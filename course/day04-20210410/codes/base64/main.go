package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	ctx := []byte("我爱中反对可怜见看来是简单快乐发动机开始了就国k/fdljfkdl1")
	rs := base64.RawStdEncoding.EncodeToString(ctx)
	fmt.Println(rs)

	txt, _ := base64.RawStdEncoding.DecodeString(rs)

	fmt.Println(string(txt))

	rs = base64.StdEncoding.EncodeToString(ctx)
	fmt.Println(rs)

	txt, _ = base64.StdEncoding.DecodeString(rs)

	fmt.Println(string(txt))

	rs = base64.RawURLEncoding.EncodeToString(ctx)
	fmt.Println(rs)

	txt, _ = base64.RawURLEncoding.DecodeString(rs)

	fmt.Println(string(txt))

	rs = base64.URLEncoding.EncodeToString(ctx)
	fmt.Println(rs)

	txt, _ = base64.URLEncoding.DecodeString(rs)

	fmt.Println(string(txt))
}
