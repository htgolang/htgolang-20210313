package persistent

import (
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"io"
	"strconv"
	"usermanager/model"
)

type CsvUserPersistent struct {
	Rw io.ReadWriter
}

func NewCsvUserPersistent(rw io.ReadWriter) *CsvUserPersistent {
	return &CsvUserPersistent{rw}
}

func (c *CsvUserPersistent) Encode(users map[int]*model.User) error {
	writer := csv.NewWriter(c.Rw)
	for _, user := range users {
		err := writer.Write([]string{strconv.Itoa(user.Id), user.Name, user.Addr, user.Tel, user.Password})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CsvUserPersistent) Decode(users *map[int]*model.User) error {
	reader := csv.NewReader(c.Rw)
	for {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		(*users)[id] = &model.User{
			Id:       id,
			Name:     record[1],
			Addr:     record[2],
			Tel:      record[3],
			Password: record[4],
		}
	}
}

type GobUserPersistent struct {
	Rw io.ReadWriter
}

func NewGobUserPersistent(rw io.ReadWriter) *GobUserPersistent {
	return &GobUserPersistent{rw}
}

func (g *GobUserPersistent) Encode(users map[int]*model.User) error {
	encoder := gob.NewEncoder(g.Rw)
	err := encoder.Encode(users)
	return err
}

func (g *GobUserPersistent) Decode(users *map[int]*model.User) error {
	decoder := gob.NewDecoder(g.Rw)
	err := decoder.Decode(users)
	return err
}

type JsonUserPersistent struct {
	Rw io.ReadWriter
}

func NewJsonUserPersistent(rw io.ReadWriter) *JsonUserPersistent {
	return &JsonUserPersistent{rw}
}

func (j *JsonUserPersistent) Encode(users map[int]*model.User) error {
	encoder := json.NewEncoder(j.Rw)
	err := encoder.Encode(model.Users)
	return err
}

func (j *JsonUserPersistent) Decode(users *map[int]*model.User) error {
	decoder := json.NewDecoder(j.Rw)
	err := decoder.Decode(&model.Users)
	return err
}
