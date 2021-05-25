package user

import (
	"fmt"
)

func Pre(t ,id, username, email, phone, password string) (PreType, error) {
	if t == "gob"{
		return NewCsv(id, username, email, phone, password) ,nil
	} 
	if t == "csv"{
		return NewGob(id, username, email, phone, password) ,nil
	}
	if t == "json"{
		return NewJson(id, username, email, phone, password) ,nil
	}
	return nil, fmt.Errorf("%s","Wrong persistence type pass")
	
}