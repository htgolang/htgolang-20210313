package usermanage

import (
	"encoding/csv"
	"encoding/gob"
	"strconv"
	"os"
	"user/module"
)

type GobPersisten struct{}

func (g GobPersisten) Save(users []module.User, path string) error {
	// encoder := gob.NewEncoder()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	gob.Register([]module.User{})

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {

		return err
	}
	return nil
}

func (g GobPersisten) Load(path string) ([]module.User, error) {
	var users []module.User

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil

}

type CsvPersistent struct{}

func (g CsvPersistent) Save(users []module.User, path string) error {
	//把[]Users转换为[][]string
	var allUser [][]string = [][]string{{"ID", "Name", "Addr", "Tel", "Passwd"}}
	for _, u := range users {
		allUser = append(allUser, []string{strconv.Itoa(u.Id), u.Name, u.Addr, u.Tel, u.Passwd})
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(allUser)
	if err != nil {
		return err
	}
	return nil
}

func (c CsvPersistent) Load(path string) ([]module.User, error) {
	var users []module.User

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	allUser, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for i := 1; i < len(allUser); i++ {
		user := allUser[i]
		id, _ := strconv.Atoi(user[0])
		users = append(users, module.User{id, user[1], user[2], user[3], user[4]})
	}

	return users, nil

}
