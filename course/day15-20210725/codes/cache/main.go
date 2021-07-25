package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/client/cache"
)

func main() {
	// cachePool, err := cache.NewCache("memory", `{"interval":60}`)
	cachePool, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	// 放 ttl
	// 判断存在不存在
	// 拿
	// 删
	if err != nil {
		log.Fatal(err)
	}

	value, err := cachePool.Get(context.TODO(), "name")
	if err == nil {
		if name, ok := value.(string); ok {
			fmt.Println("get:", name)
		}
	} else {
		fmt.Println("get:", err)
	}

	err = cachePool.Put(context.TODO(), "name", "kk", 10*time.Second)
	fmt.Println("put: ", err)

	value, err = cachePool.Get(context.TODO(), "name")
	if err == nil {
		if name, ok := value.(string); ok {
			fmt.Println("get: ", name)
		}
	} else {
		fmt.Println("get: ", err)
	}

	// err = cachePool.Delete(context.TODO(), "name")
	// fmt.Println("delete:", err)
	// time.Sleep(12 * time.Second)

	// value, err = cachePool.Get(context.TODO(), "name")
	// if err == nil {
	// 	if name, ok := value.(string); ok {
	// 		fmt.Println("get: ", name)
	// 	}
	// } else {
	// 	fmt.Println("get: ", err)
	// }
}
