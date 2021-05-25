package service

import (
	"errors"
	"os"
	"path/filepath"
	"usermanager/model"
	"usermanager/persistent"
)

type DbService struct {
	Type string
}

var Db DbService

func (d DbService) EncodeUser() error {
	var p persistent.UserPersistent
	switch d.Type {
	case "gob":
		path := filepath.Join("db", "user.gob")
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		p = persistent.NewCsvUserPersistent(f)
	case "csv":
		path := filepath.Join("db", "user.csv")
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		p = persistent.NewGobUserPersistent(f)
	case "json":
		path := filepath.Join("db", "user.json")
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		p = persistent.NewJsonUserPersistent(f)
	default:
		return errors.New("未知类型!")
	}
	err := p.Encode(model.Users)
	return err
}

func (d DbService) DecodeUser() error {
	var p persistent.UserPersistent
	switch d.Type {
	case "gob":
		path := filepath.Join("db", "user.gob")
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		p = persistent.NewGobUserPersistent(f)
	case "csv":
		path := filepath.Join("db", "user.csv")
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		p = persistent.NewCsvUserPersistent(f)
	case "json":
		path := filepath.Join("db", "user.json")
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		p = persistent.NewJsonUserPersistent(f)
	default:
		return errors.New("未知类型!")
	}
	err := p.Decode(&model.Users)
	return err
}
