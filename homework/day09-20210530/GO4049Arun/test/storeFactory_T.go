package main

import (
	"encoding/json"
	"fmt"
)

type AlipayRemoteReqStruct struct {
	Ono string `json:ono`
	OrderItem []*Item `json:item`
}
type Item struct {
	Ono string `json:ono`
	Oid int    `json:oid`
}

func main() {
	var m AlipayRemoteReqStruct
	m.Ono = "12345"
	m.OrderItem = append(m.OrderItem, &Item{ "Shanghai_VPN",  1})
	m.OrderItem = append(m.OrderItem, &Item{"Beijing_VPN", 2})
	fmt.Println(m)
	bytes, _ := json.Marshal(m)
	fmt.Printf("json:%s\n", bytes)
}