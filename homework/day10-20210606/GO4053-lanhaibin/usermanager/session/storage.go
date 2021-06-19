package session

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const store_dir = "temp/session"

func DumpSession(s *Session) error {
	store_path := filepath.Join(store_dir, s.Id)
	f, err := os.OpenFile(store_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	return encoder.Encode(s)
}

func LoadSession(id string) (Session, error) {
	s := Session{}

	store_path := filepath.Join(store_dir, id)
	f, err := os.Open(store_path)
	if err != nil {
		return s, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&s)
	return s, err
}

func LoadOrNewSession(id string) Session {
	s, err := LoadSession(id)
	if err != nil {
		s = NewSession()
	}
	return s
}

func DeleteSession(id string) {
	store_path := filepath.Join(store_dir, id)
	os.Remove(store_path)
}
