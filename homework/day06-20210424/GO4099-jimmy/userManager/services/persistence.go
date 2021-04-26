package services

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"userManager/model"
	"userManager/utils"
)

const (
	gobPath = "user.gob"
	csvPath = "user.csv"
)

type PersistenceService struct{}

var P = model.Persistence{Types: "csv"}

func (PersistenceService) Pst() {
START:
	p := utils.Input("请选择持久化方式(gob/csv): ")
	if p == "gob" {
		P = model.Persistence{Types: p}
	} else if p == "csv" {
		P = model.Persistence{Types: p}
	} else {
		fmt.Println("输入错误")
		goto START
	}
}

func decode() (users map[int]model.User) {
	var err error
	if P.Types == "gob" {
		users, err = P.Decode(GobDecode)
		if err != nil {
			log.Fatal(err)
		}
	} else if P.Types == "csv" {
		users, err = P.Decode(CsvDecode)
		if err != nil {
			log.Fatal(err)
		}
	}
	return users
}

func encode(users map[int]model.User) {
	if P.Types == "gob" {
		err := P.Encode(GobEncode, users)
		if err != nil {
			log.Fatal(err)
		}
	} else if P.Types == "csv" {
		err := P.Encode(CsvEncode, users)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GobEncode(user map[int]model.User) error {
	file, err := os.Create(gobPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewEncoder(file)
	return decoder.Encode(user)
}

func GobDecode() (map[int]model.User, error) {
	users := make(map[int]model.User)
	file, err := os.Open(gobPath)
	if err != nil {
		if os.IsNotExist(err) {
			return users, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func CsvDecode() (map[int]model.User, error) {
	user := make(map[int]model.User)
	file, err := os.Open(csvPath)
	if err != nil {
		if os.IsNotExist(err) {
			return user, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		id, _ := strconv.Atoi(line[0])
		user[id] = model.User{id, line[1], line[2], line[3], line[4]}
	}
	return user, nil
}

func CsvEncode(user map[int]model.User) error {
	file, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, v := range user {
		writer.Write([]string{strconv.Itoa(v.Id), v.Name, v.Addr, v.Tel, v.Password})
	}
	writer.Flush()
	return nil
}
