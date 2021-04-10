package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandText(n int) string {
	rs := make([]byte, n)
	for i := 0; i < n; i++ {
		rs[i] = chars[rand.Intn(len(chars))]
	}
	return string(rs)
}

func Int2StrV1(n int) string {
	return strconv.Itoa(n)
}

func Int2StrV2(n int) string {
	return fmt.Sprintf("%d", n)
}
