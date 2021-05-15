package users

import (
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

//define the interface
type Persistence interface {
	Load() (map[int]User, error) //return
	Store(map[int]User) error    //param
}

//json format
type JsonFile struct {
	name string
}

func NewJsonFile(name string) JsonFile {
	return JsonFile{name + ".json"}
}

func (j JsonFile) Load() (map[int]User, error) {
	//read file
	bytes, err := ioutil.ReadFile(j.name)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]User), nil
		}
		return nil, err
	}
	var users map[int]User
	err = json.Unmarshal(bytes, &users)
	return users, err
}

func (j JsonFile) Store(users map[int]User) error {
	//rename the file before save
	if _, err := os.Stat(j.name); err == nil {
		os.Rename(j.name, strconv.FormatInt(time.Now().Unix(), 10)+".user.json")
	}

	//keep three files
	if names, err := filepath.Glob("*.user.json"); err == nil {
		if len(names) == 4 {
			sort.Sort(sort.Reverse(sort.StringSlice(names)))
			for _, name := range names[3:] {
				os.Remove(name)
			}
		}
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.name, bytes, os.ModePerm)
}

//gob format
type GobFile struct {
	name string
}

func NewGobFile(name string) GobFile {
	return GobFile{name + ".gob"}
}

func (g GobFile) Load() (map[int]User, error) {
	file, err := os.Open(g.name)
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		}
		return nil, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	var users map[int]User
	err = decoder.Decode(&users)
	return users, err
}

func (g GobFile) Store(users map[int]User) error {
	//rename the file before save
	if _, err := os.Stat(g.name); err == nil {
		os.Rename(g.name, strconv.FormatInt(time.Now().Unix(), 10)+".user.gob")
	}

	//keep three files
	if names, err := filepath.Glob("*.user.gob"); err == nil {
		if len(names) == 4 {
			sort.Sort(sort.Reverse(sort.StringSlice(names)))
			for _, name := range names[3:] {
				os.Remove(name)
			}
		}
	}

	file, err := os.Create(g.name)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(users)
}
