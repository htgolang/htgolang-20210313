package utils

import (
	"encoding/gob"
	"os"
	"userManager/model"
)

const (
	gobPath = "user.gob"
	csvPath = "user.csv"
)

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

func CsvDecode() (map[int]model.User, error) {}
func CsvEncode(user map[int]model.User) error {}
