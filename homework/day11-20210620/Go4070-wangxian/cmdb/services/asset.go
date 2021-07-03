package services

import (
	"cmdb/db"
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"log"
)

func QueryAsset() []models.Asset {
	var ats []models.Asset
	sql := `select id,ip,addr from asset;`
	rows, err := db.DbConn.Query(sql)
	if err != nil {
		log.Printf("query asset from db error. %s", err)
		return nil
	}
	for rows.Next() {
		var at models.Asset
		err := rows.Scan(&at.Id, &at.Ip, &at.Addr)
		if err != nil {
			log.Printf("query asset: scan data from rows error. %s", err)
			continue
		}
		ats = append(ats, at)
	}
	return ats
}

func DleteAsset(id int64) {
	sql := `delete from asset where id=?;`
	db.DbConn.Exec(sql, id)
}

func AddAsset(ip, addr string) error {
	if ip == "" || addr == "" {
		return fmt.Errorf("ip或位置不能为空")
	}

	if utils.CheckAssetExists(-1, ip) {
		return fmt.Errorf("IP冲突")
	}

	sql := `insert into asset(ip,addr,create_at,update_at) values(?,?,now(),now());`
	_, err := db.DbConn.Exec(sql, ip, addr)
	if err != nil {
		return fmt.Errorf("添加资产,请重试")
	}
	return nil
}

func EditAsset(id int64, ip, addr string) error {
	if ip == "" || addr == "" {
		return fmt.Errorf("ip或位置不能为空")
	}

	if utils.CheckAssetExists(id, ip) {
		return fmt.Errorf("IP冲突")
	}

	sql := `update asset set ip=?,addr=? where id=?;`
	_, err := db.DbConn.Exec(sql, ip, addr, id)
	if err != nil {
		return fmt.Errorf("资产添加失败,请重试")
	}
	return nil
}

func QueryAssetByID(id int64) *models.Asset {
	var asset models.Asset
	sql := `select id,ip,addr from asset where id=?;`
	err := db.DbConn.QueryRow(sql, id).Scan(&asset.Id, &asset.Ip, &asset.Addr)
	if err != nil {
		log.Printf("query_asset_byid: scan data from row error. %s", err)
		return nil
	}
	return &asset
}
