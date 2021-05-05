package service

import (
	"encoding/csv"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"usermanager/model"
)

type DbService struct {
	Type string
}

var Db DbService

func (d DbService) EncodeUser() error {
	switch d.Type {
	case "gob":
		path := filepath.Join("db", "user.gob")
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		encoder := gob.NewEncoder(f)
		err = encoder.Encode(model.Users)
		if err != nil {
			return err
		}
		return nil
	case "csv":
		path := filepath.Join("db", "user.csv")
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		encoder := csv.NewWriter(f)
		defer encoder.Flush()
		for _, user := range model.Users {
			encoder.Write([]string{strconv.Itoa(user.Id), user.Name, user.Addr, user.Tel, user.Password})
		}
		return nil
	}
	return errors.New("未知类型!")
}

func (d DbService) DecodeUser() error {
	switch d.Type {
	case "gob":
		path := filepath.Join("db", "user.gob")
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		decoder := gob.NewDecoder(f)
		err = decoder.Decode(&model.Users)
		// fmt.Println(model.Users[1].Password)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "csv":
		path := filepath.Join("db", "user.csv")
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		decoder := csv.NewReader(f)
		lines, err := decoder.ReadAll()
		if err != nil {
			return err
		}
		model.Users = make(map[int]*model.User)
		for _, line := range lines {
			id, _ := strconv.Atoi(line[0])
			model.Users[id] = &model.User{
				Id:       id,
				Name:     line[1],
				Addr:     line[2],
				Tel:      line[3],
				Password: line[4],
			}
		}
		return nil
	}
	return nil
}
