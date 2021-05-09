package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if date, err := time.Parse("2006-01-02", s); err != nil {
		return err
	} else {
		*d = Date(date)
	}
	return nil
}

type Date2 time.Time

func (d Date2) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format("2006-01-02")), nil
}

func (d *Date2) UnmarshalText(b []byte) error {
	s := string(b)
	if date, err := time.Parse("2006-01-02", s); err != nil {
		return err
	} else {
		*d = Date2(date)
	}
	return nil
}

type User struct {
	Id        int
	Name      string
	Birthday  time.Time
	Birthday2 Date
	Birthday3 Date2
}

func main() {

	u := User{1, "", time.Now(), Date(time.Now()), Date2(time.Now())}

	b, err := json.MarshalIndent(u, "", "\t")
	fmt.Println(err, string(b))

	txt := `
	{
        "Id": 2,
        "Name": "kk",
        "Birthday": "2021-02-09T15:42:07.9917883+08:00",
        "Birthday2": "2021-03-09",
        "Birthday3": "2021-03-09"
	}
	`

	var us User
	err = json.Unmarshal([]byte(txt), &us)
	fmt.Println(err, us)
}
