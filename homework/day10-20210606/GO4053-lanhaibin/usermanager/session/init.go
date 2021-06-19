package session

import "os"

func init() {
	os.MkdirAll(store_dir, os.ModePerm)
}
