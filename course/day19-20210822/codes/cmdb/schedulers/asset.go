package schedulers

import (
	"cmdb/forms"
	"cmdb/services"
	"cmdb/utils"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

type assetScheduler struct {
	locker   sync.RWMutex
	enable   bool
	interval time.Duration
	ipRanges []*utils.IPRange
	ports    []*utils.PortRange
}

func newAssetScheduler() *assetScheduler {
	return &assetScheduler{}
}

var AssetScheduler = newAssetScheduler()

func (s *assetScheduler) SyncConfig() {
	s.locker.Lock()
	defer s.locker.Unlock()

	form := new(forms.AssetScanConfigForm)
	config := services.ConfigService.GetAssetScan()
	form.FillModel(config)

	s.enable = form.Enable
	s.interval = time.Duration(form.Interval) * time.Second
	s.ipRanges = utils.ParseIPRange(form.Addr)
	s.ports = utils.ParsePorts(form.Ports)

}

func (s *assetScheduler) Start() {
	s.SyncConfig()
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		prevTime := time.Unix(0, 0)

		for {
			// 检查是否需要运行
			// 记录上次运行的时间
			// now  >= prev + interval
			now := time.Now()
			s.locker.RLock()
			logs.Debug("check sync asset: %t, now: %v, prevTime: %v", s.enable, now, prevTime)
			if s.enable && prevTime.Add(s.interval).Before(now) {
				logs.Info("sync asset info: %#v, %#v", s.ipRanges, s.ports)
				services.AssetService.Scan(s.ipRanges, s.ports)
				prevTime = now
			}
			s.locker.RUnlock()
			<-ticker.C
		}
	}()
}
