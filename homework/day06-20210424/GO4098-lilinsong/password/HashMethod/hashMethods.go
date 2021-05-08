package HashMethod

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

const readBuf = 1024 * 1024 * 1024

type Methods interface {
	Md5([]byte, string) string
	Sha1([]byte, string) string
	Sha256([]byte, string) string
	Sha512([]byte, string) string
	method(string, string, []byte) string
	FileMethod() FileMethod
	pwdMethod() pwdMethod
}

type HashEncryption struct {
	startCha chan struct{}
	stopCha  chan struct{}
	wg       sync.WaitGroup
}

func (h *HashEncryption) fileHash(file string) string {
	hashed := md5.New()
	fileObj, _ := os.Open(file)
	defer fileObj.Close()
	buffer := make([]byte, readBuf)
	if _, err := io.CopyBuffer(hashed, fileObj, buffer); err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(hashed.Sum(nil))
}

func (h *HashEncryption) Md5(text []byte, file string) string {
	defer h.wg.Done()
	hashed := md5.New()
	if file != "" {
		return h.fileHash(file)
	} else {

	END:
		for {
			select {
			case <-h.startCha:
				if _, err := hashed.Write(text); err != nil {
					fmt.Println(err)
					return ""
				}
			case <-h.stopCha:
				break END
			}
		}
	}
	return hex.EncodeToString(hashed.Sum(nil))
}

func (h *HashEncryption) Sha1(text []byte, file string) string {
	defer h.wg.Done()
	hashed := md5.New()
	if file == "" {
		h.fileHash(file)
		return hex.EncodeToString(hashed.Sum(nil))
	} else {

	END:
		for {
			select {
			case <-h.startCha:
				if _, err := hashed.Write(text); err != nil {
					fmt.Println(err)
					return ""
				}
			case <-h.stopCha:
				break END
			}
		}
	}
	return hex.EncodeToString(hashed.Sum(nil))
}
func (h *HashEncryption) Sha256(text []byte, file string) string {
	defer h.wg.Done()
	hashed := md5.New()
	if file == "" {
		h.fileHash(file)
		return hex.EncodeToString(hashed.Sum(nil))
	} else {

	END:
		for {
			select {
			case <-h.startCha:
				if _, err := hashed.Write(text); err != nil {
					fmt.Println(err)
					return ""
				}
			case <-h.stopCha:
				break END
			}
		}
	}
	return hex.EncodeToString(hashed.Sum(nil))
}
func (h *HashEncryption) Sha512(text []byte, file string) string {
	hashed := sha512.New()
	defer h.wg.Done()
END:
	for {
		select {
		case <-h.startCha:
			if _, err := hashed.Write(text); err != nil {
				fmt.Println(err)
				return ""
			}
		case <-h.stopCha:
			break END
		}

	}
	return hex.EncodeToString(hashed.Sum(nil))
}

func (h *HashEncryption) method(method string, file string, hash []byte) string {
	switch method {
	case "sha1":
		return h.Sha1(hash, file)
	case "sha256":
		return h.Sha256(hash, file)
	case "sha512":
		return h.Sha512(hash, file)
	case "md5":
		return h.Md5(hash, file)
	default:
		fmt.Println("没有匹配到加密方法, 默认使用 md5 方法, 是否继续(y/n)")
		var tmp string
		fmt.Print("请输入: ")
		fmt.Scan(&tmp)
		if strings.ToLower(tmp) != "y" {
			return ""
		}
		return h.Md5(hash, file)
	}
}
