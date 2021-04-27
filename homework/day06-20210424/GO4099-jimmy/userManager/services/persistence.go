package services

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"userManager/model"
	"userManager/utils"
)

type PersistenceService struct{}

const confPath = "./conf/persistence.conf"

var P = model.Persistence{
	Types:   readConf(),
	GobPath: "user.gob",
	CsvPath: "user.csv",
	Count:   3,
}

func (PersistenceService) Pst() {
	fmt.Println("现在的持久化方式是", P.Types)
START:
	p := utils.Input("请选择持久化方式(gob/csv): ")
	switch p {
	case "gob", "csv":
		if p != P.Types {
			users := Decode()
			P.Types = p
			err := ioutil.WriteFile(confPath, []byte("persistence "+p), 0644)
			if err != nil {
				log.Fatal(err)
			}
			Encode(users)
		}
	default:
		fmt.Println("输入错误")
		goto START
	}
}

func readConf() string {
	bytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(bytes), " ")[1]
}

func Decode() (users map[int]model.User) {
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

func Encode(users map[int]model.User) {
	logSplit()
	err := P.Encode(GobEncode, users)
	if err != nil {
		log.Fatal(err)
	}
	err = P.Encode(CsvEncode, users)
	if err != nil {
		log.Fatal(err)
	}
}
func logSplit() {
	for _, path := range []string{P.CsvPath, P.GobPath} {
		f, err := os.Open(path)
		if os.IsNotExist(err) {
			continue
		}
		_ = f.Close()
		err = os.Rename(path, path+"_"+strconv.FormatInt(time.Now().Unix(), 10))
		if err != nil {
			log.Fatal(err)
		}
		files, err := filepath.Glob(path + "_*")
		if err != nil {
			log.Fatal(err)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(files)))
		if len(files) <= P.Count {
			continue
		}
		for _, file := range files[P.Count:] {
			err := os.Remove(file)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func GobEncode(user map[int]model.User) error {
	file, err := os.Create(P.GobPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewEncoder(file)
	return decoder.Encode(user)
}

func GobDecode() (map[int]model.User, error) {
	users := make(map[int]model.User)
	file, err := os.Open(P.GobPath)
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
	file, err := os.Open(P.CsvPath)
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
	file, err := os.Create(P.CsvPath)
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
