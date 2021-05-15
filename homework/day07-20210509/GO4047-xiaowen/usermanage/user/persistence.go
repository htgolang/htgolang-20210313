package user

import (
	"path/filepath"
	"os"
	"encoding/gob"
	"encoding/csv"
	"encoding/json"
	"log"
	"strconv"
)
type Presistnece struct{
	Id string 
	Username string
	Email string
	Phone string
	Password string
}
func (p *Presistnece) CsvPre(id, username, email, phone, password string){
	dinfo,err := filepath.Glob("D:\\go\\src\\go-dev\\usermanage\\user\\*.csv")
	if err != nil{
		log.Fatalln(err)
	}
	if len(dinfo) == 0 {
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\1.csv")
		defer file.Close()
		if cerr != nil{
			log.Fatalln(cerr)
		}
		write := csv.NewWriter(file)
		werr := write.Write([]string{id, username, email, phone, password})
		write.Flush()
		if werr != nil{
			log.Fatalln(werr)
		}
	} else if len(dinfo) == 4{
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".csv")
		defer file.Close()
		reer := os.Remove(dinfo[0])
		if reer != nil{
			log.Fatalln(reer)
		}
		if cerr != nil{
			log.Fatal(cerr)
		}
		write := csv.NewWriter(file)
		werr := write.Write([]string{id, username, email, phone, password})
		write.Flush()
		if werr != nil{
			log.Fatalln(werr)
		}
	} else {
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".csv")
		defer file.Close()
		if cerr != nil{
			log.Fatal(cerr)
		}
		write := csv.NewWriter(file)
		werr := write.Write([]string{id, username, email, phone, password})
		write.Flush()
		if werr != nil{
			log.Fatalln(werr)
		}
	}
}

func (p *Presistnece) GobPre(id, username, email, phone, password string){
	dinfo,err := filepath.Glob("D:\\go\\src\\go-dev\\usermanage\\user\\*.gob")
	if err != nil{
		log.Fatalln(err)
	}
	if len(dinfo) == 0 {
		gob.Register(&Useradmin{})
		users := []*Useradmin{
			{id, username, email, phone, password},
		}
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\1.gob")
		defer file.Close()
		if cerr != nil{
			log.Fatalln(cerr)
		}
		encoder := gob.NewEncoder(file)
		eerr := encoder.Encode(users)
		if eerr != nil{
			log.Fatalln(eerr)
		}
	} else if len(dinfo) == 4 {
		gob.Register(&Useradmin{})
		users := []*Useradmin{
			{id, username, email, phone, password},
		}
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".gob")
		defer file.Close()
		reer := os.Remove(dinfo[0])
		if reer != nil{
			log.Fatalln(reer)
		}
		if cerr != nil{
			log.Fatal(cerr)
		}
		encoder := gob.NewEncoder(file)
		eerr := encoder.Encode(users)
		if eerr != nil{
			log.Fatalln(eerr)
		}
	} else{
		gob.Register(&Useradmin{})
		users := []*Useradmin{
			{id, username, email, phone, password},
		}
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".gob")
		defer file.Close()
		if cerr != nil{
			log.Fatal(cerr)
		}
		encoder := gob.NewEncoder(file)
		eerr := encoder.Encode(users)
		if eerr != nil{
			log.Fatalln(eerr)
		}
	}

}



func (p *Presistnece) JsonPre(id, username, email, phone, password string){
	dinfo,err := filepath.Glob("D:\\go\\src\\go-dev\\usermanage\\user\\*.json")
	if err != nil{
		log.Fatalln(err)
	}
	if len(dinfo) == 0 {
		users := Presistnece{id, username, email, phone, password}
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\1.json")
		defer file.Close()
		if cerr != nil{
			log.Fatalln(cerr)
		}
		encoder := json.NewEncoder(file)
		eerr := encoder.Encode(users)
		if eerr != nil{
			log.Fatalln(eerr)
		}
	} else if len(dinfo) == 4 {
		users := Presistnece{id, username, email, phone, password}
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".json")
		defer file.Close()
		reer := os.Remove(dinfo[0])
		if reer != nil{
			log.Fatalln(reer)
		}
		if cerr != nil{
			log.Fatal(cerr)
		}
		encoder := json.NewEncoder(file)
		eerr := encoder.Encode(&users)
		if eerr != nil{
			log.Fatalln(eerr)
		}
	} else{
		users := Presistnece{id, username, email, phone, password}
		file, cerr := os.Create("D:\\go\\src\\go-dev\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".json")
		defer file.Close()
		if cerr != nil{
			log.Fatal(cerr)
		}
		encoder := json.NewEncoder(file)
		eerr := encoder.Encode(&users)
		if eerr != nil{
			log.Fatalln(eerr)
		}
	}
}

// type preDtat struct{}

// func (t preDtat) Gob(){}
// func (t preDtat) Csv(){}
// func (t preDtat) Json(){}

