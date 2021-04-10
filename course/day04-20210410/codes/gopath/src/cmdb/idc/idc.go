package idc

import (
	"fmt"
	"prometheus/agent"
)

func Add() {
	fmt.Println("idc add")
	agent.Add()
}
