package services

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"

	"cmdb/models"
	"cmdb/utils"
)

type assetService struct {
	locker      sync.RWMutex
	concurrency int
	isScanning  bool
}

func newAssetService() *assetService {
	return &assetService{
		isScanning:  false,
		concurrency: 10,
		locker:      sync.RWMutex{},
	}
}

var AssetService = newAssetService()

func (s *assetService) Query(q string) []models.Asset {
	ormer := orm.NewOrm()

	queryset := ormer.QueryTable(new(models.Asset))
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)
	if q != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("Addr__icontains", q)
		queryCond = queryCond.Or("Description__icontains", q)
		cond = cond.AndCond(queryCond)
	}
	var assets []models.Asset

	queryset.SetCond(cond).All(&assets)

	return assets
}

func (s *assetService) IsScanning() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()
	return s.isScanning
}

func (s *assetService) updateStatus(scanning bool) {
	s.locker.Lock()
	s.isScanning = scanning
	s.locker.Unlock()
}

func (s *assetService) Scan(ipRanges []*utils.IPRange, portRanges []*utils.PortRange) error {
	if s.IsScanning() {
		return nil
	}
	s.updateStatus(true)
	limit := s.concurrency
	go func() {
		var wg sync.WaitGroup
		incr := 0
		starTime := time.Now()
		for _, ipRange := range ipRanges {
			for ip := ipRange.Start(); ip <= ipRange.End(); ip++ {
				host := utils.Int2Ip(ip)
				incr++
				wg.Add(1)

				go func() {
					for _, portRange := range portRanges {
						for port := portRange.Start(); port <= portRange.End(); port++ {
							ok := utils.TestConnect(net.JoinHostPort(host, strconv.Itoa(port)))
							logs.Debug("test connect: %s[%d](%t)", host, port, ok)
							if ok {
								// 更新数据库
								s.syncUpdate(host)
								break
							}
						}
					}
					wg.Done()
				}()

				if incr >= limit {
					wg.Wait()
					incr = 0
				}
			}
		}

		if incr > 0 {
			wg.Wait()
		}

		s.updateStatus(false)
		logs.Info("sync asset elasped: %v", time.Now().Sub(starTime))
	}()
	return nil
}

func (s *assetService) syncUpdate(addr string) {
	ormer := orm.NewOrm()
	asset := &models.Asset{Addr: addr}
	if err := ormer.Read(asset, "Addr"); err != nil {
		asset.Status = 0
		ormer.Insert(asset)
	} else if asset.DeletedAt != nil {
		asset.Status = 0
		asset.DeletedAt = nil
		ormer.Update(asset)
	}

}
