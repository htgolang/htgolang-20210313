package users

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

const (
	//The max number of times for the password attempts
	MathAuth = 3
	// password = "202cb962ac59075b964b07152d234b70" //md5 32bit result
	//Store it to file
	passwordFile = ".password"
	userFile     = "user"
)

//define the struct
type User struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Birth time.Time `json:"birthday"`
	Tel   string    `json:"tel"`
	Addr  string    `json: "addr"`
	Desc  string    `json: "desc`
}

//optimize the print function
func (u User) String() string {
	return fmt.Sprintf("ID: %d\nName:%s\nBirth:%s\nAddr:%s\nPhone:%s\nDesc:%s\n", u.ID, u.Name, u.Birth.Format("2006-01-02"), u.Addr, u.Tel, u.Desc)
}

//redefine user
//no need to pass the user arg to the the functions
// User ==> map[string]string
// var users map[int]map[string]string
// var users map[int]User

// var persistence Persistence = NewGobFile(userFile)
var persistence Persistence = NewJsonFile(userFile)

/*
//gob operation
func loadUser() map[int]User {
	//decode
	users := map[int]User{}
	if file, err := os.Open(userFile); err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&users)
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("error: ", err)
		}
	}
	return users
}

func storeUser(users map[int]User) {
	//rename the file before save
	if _, err := os.Stat(userFile); err == nil {
		os.Rename(userFile, strconv.FormatInt(time.Now().Unix(), 10)+".user.gob")
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
	//encode
	if file, err := os.Create(userFile); err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(users)
	}
}
*/

// func init() {
// 	users = make(map[int]User)
// }

//public input function
func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input) //remove the space
}

//verify the password
func Auth() bool {
	//read password from the file
	password, err := ioutil.ReadFile(passwordFile)
	if err == nil && len(password) > 0 {
		//check the password
		//if the password passed, return true
		// var input string
		for i := 0; i < MathAuth; i++ {
			fmt.Print("Pls enter a password: ")
			// fmt.Scan(&input)
			bytes, _ := gopass.GetPasswd()
			// from string to []byte,
			pass := fmt.Sprintf("%x", md5.Sum(bytes))
			//password is a byte
			if string(password) == pass {
				//success
				return true
			} else {
				//fail
				fmt.Println("Wrong password!")
			}
		}
		return false //failed three times, will be false
	} else {
		if len(password) == 0 || os.IsNotExist(err) {
			//check the password file, not exist then init password
			fmt.Print("Pls provide the init password: ")
			bytes, _ := gopass.GetPasswd()
			//check md5 value, and write to the file
			ioutil.WriteFile(passwordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			return true
		} else {
			//other reasons
			fmt.Println("error", err)
			return false
		}
	}
}

//format the output
// func printUser(pu int, user map[string]string) {
/*
func printUser(user User) {
	fmt.Println("ID: ", pu)
	fmt.Println("Name: ", user["name"])
	fmt.Println("Birth: ", user["birth"])
	fmt.Println("Tel: ", user["tel"])
	fmt.Println("Addr: ", user["addr"])
	fmt.Println("Desc: ", user["desc"])

	fmt.Println("ID: ", user.ID)
	fmt.Println("Name: ", user.Name)
	//format: from time to string
	fmt.Println("Birth: ", user.Birth.Format("2021-01-01"))
	fmt.Println("Tel: ", user.Tel)
	fmt.Println("Addr: ", user.Addr)
	fmt.Println("Desc: ", user.Desc)
}
*/

//query user function
func QueryUser() {
	q := InputString("Pls enter a query string: ")
	fmt.Println("==========================")
	// for k, v := range users {
	// users := loadUser()
	users, err := persistence.Load()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, v := range users {
		//if contain the related string, then print
		// if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["add"], q) || strings.Contains(v["desc"], q) {
		if strings.Contains(v.Name, q) || strings.Contains(v.Tel, q) || strings.Contains(v.Addr, q) || strings.Contains(v.Desc, q) {
			// printUser(k, v)
			fmt.Println(v)
			fmt.Println("---------------------------")
		}
	}
	fmt.Println("==========================")
}

//public input user info
// func inputUser() map[string]string {
func inputUser(id int) User {
	//user := map[string]string{}
	var user User
	user.ID = id
	user.Name = InputString("Pls enter a name: ")
	//from string to time
	// user.Birth = InputString("Pls enter the birthday: ")
	bir, _ := time.Parse("2021-04-22", InputString("Pls enter the birthday: "))
	user.Birth = bir
	user.Tel = InputString("Pls enter the phone number: ")
	user.Addr = InputString("Pls enter an address: ")
	user.Desc = InputString("Pls enter the desc: ")
	return user //return struct
}

//get the user id
func getId() (int, error) {
	var id int
	// users := loadUser()
	users, err := persistence.Load()
	if err != nil {
		fmt.Println("Error: ", err)
		return -1, err
	}
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1, nil
}

//add user function
func AddUser() {
	id, err := getId()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	//add
	// user := inputUser()
	user := inputUser(id)
	// users := loadUser()
	users, err := persistence.Load()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	users[id] = user
	// storeUser(users)
	if err := persistence.Store(users); err == nil {
		fmt.Println("add success.")
	} else {
		fmt.Println("add fail.")
	}
}

//modify user function
func ModifyUser() {
	idString := InputString("Pls enter the modify user id: ")
	//string to int
	if id, err := strconv.Atoi(idString); err == nil {
		//key in map or not
		// users := loadUser()
		users, err := persistence.Load()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if user, ok := users[id]; ok {
			fmt.Println("The modify info: ")
			// printUser(id, user)
			fmt.Println(user)
			input := InputString("confirm(y/n)? ")
			if input == "y" || input == "Y" {
				// user := inputUser()
				user := inputUser(id)
				users[id] = user
				// storeUser(users)
				if err := persistence.Store(users); err == nil {
					fmt.Println("modify success.")
				} else {
					fmt.Println("modify fail.")
				}
			}
		} else {
			fmt.Println("User id not exist.")
		}
	} else {
		fmt.Println("Wrong id.")
	}
}

//delete user function
func DeleteUser() {
	//similar with modify
	idString := InputString("Pls enter the delete user id: ")
	//string to int
	if id, err := strconv.Atoi(idString); err == nil {
		//key in map or not
		// users := loadUser()
		users, err := persistence.Load()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if user, ok := users[id]; ok {
			fmt.Println("The delete info: ")
			// printUser(id, user)
			fmt.Println(user)
			input := InputString("confirm(y/n)? ")
			if input == "y" || input == "Y" {
				//delete key
				delete(users, id)
				// storeUser(users)
				if err := persistence.Store(users); err == nil {
					fmt.Println("remove success.")
				} else {
					fmt.Println("remove fail.")
				}
			}
		} else {
			fmt.Println("User id not exist.")
		}
	} else {
		fmt.Println("Wrong id.")
	}
}
