package services

import "fmt"

type IDCService struct {
}

func (IDCService) Add() {
	fmt.Println("add idc")
}

func (IDCService) Modify() {
	fmt.Println("modify idc")
}

func (IDCService) Delete() {
	fmt.Println("delete idc")
}

func (IDCService) Query() {
	fmt.Println("query idc")
}
