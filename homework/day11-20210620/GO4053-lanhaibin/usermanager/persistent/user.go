package persistent

import (
	"encoding/json"
	"os"
)

type Persistent interface {
	Encode(v interface{}) error
	Decode(v interface{}) error
}

type JsonPersistent struct {
	DbPath string
}

func (j JsonPersistent) Encode(v interface{}) error {
	f, err := os.OpenFile(j.DbPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	return encoder.Encode(v)
}

func (j JsonPersistent) Decode(v interface{}) error {
	f, err := os.Open(j.DbPath)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	return decoder.Decode(v)
}
