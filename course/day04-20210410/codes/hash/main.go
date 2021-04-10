package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	ctx := []byte("我爱中国")
	hash := fmt.Sprintf("%x", md5.Sum(ctx)) // fmt.Printf("")
	fmt.Println(hash)

	hasher := md5.New()
	hasher.Write([]byte("我爱"))
	hasher.Write([]byte("中国"))
	hash = fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Println(hash)

	// sha1/sha256/sha512

	fmt.Printf("%x\n", sha1.Sum(ctx))

	hasher = sha1.New()
	hasher.Write([]byte("我爱"))
	hasher.Write([]byte("中国"))
	hash = fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Println(hash)

	fmt.Printf("%x\n", sha256.Sum256(ctx))

	hasher = sha256.New()
	hasher.Write([]byte("我爱"))
	hasher.Write([]byte("中国"))
	hash = fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Println(hash)

	fmt.Printf("%x\n", sha512.Sum512(ctx))

	hasher = sha512.New()
	hasher.Write([]byte("我爱"))
	hasher.Write([]byte("中国"))
	hash = fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Println(hash)

}
