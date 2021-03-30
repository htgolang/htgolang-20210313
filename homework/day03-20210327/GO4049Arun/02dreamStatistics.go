package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	var database = make(map[string]int)
	// read content from a file
	article, err := ioutil.ReadFile(`02IhaveADream.txt`)
	if err != nil {
		panic(err)
	}
	//(1)Main function realization
	for _, ch := range article {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			database[string(ch)]++
		}
	}

	//(2)Reverse the key and value of the database
	var databaseRev = make(map[int]string)
	for k, v := range database {
		databaseRev[v] = k
	}
	//(3)sort the key
	keys := make([]int, 0, len(databaseRev))
	for k := range databaseRev {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//(4)get the key and value that are reversed again
	for i := len(keys) - 1; i >= len(keys)-10; i-- {
		fmt.Printf("%v=>%v\n", databaseRev[keys[i]], keys[i])
	}
	//(5)此程序暂未考虑key重复的情况
}
