package persistence

import (
	"encoding/json"
	"fmt"
	"os"
)


type PerType interface{
	Encode(data interface{}) error  
	Decode(data interface{}) error 
}


type Persistence struct{
	StroagePath string
}

func(P Persistence) Encode(data interface{}) error{
	file, err := os.OpenFile(P.StroagePath,os.O_CREATE|os.O_RDWR|os.O_TRUNC,0644)
	if err != nil{
		fmt.Println(err)
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func(P Persistence) Decode(data interface{}) error{
	file, err := os.Open(P.StroagePath)
	if err != nil{
		fmt.Println(err)
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}


