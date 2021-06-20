package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pwd := "kk@123"

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	fmt.Println(string(hash), err)

	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("kk")))
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("kk1")))
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("kk2")))
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("kk@123")))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte("$2a$10$mLJdhrPlkJGVQFgmTSvvHOnrw2LV6yi2JgtoznOtclZ3eCtS7QfO."), []byte("kk@123")))
}
