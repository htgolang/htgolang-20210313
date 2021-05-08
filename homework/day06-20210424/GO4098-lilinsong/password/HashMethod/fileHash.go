package HashMethod

import (
	"fmt"
	"os"
	"sync"
)

type FileMethod interface {
	FileMd5(file, t string, wg *sync.WaitGroup) (string, error)
}

type FileHashMethod interface {
	FileMethod() FileMethod
}

type fileHash struct {
	file string
	hash *HashEncryption
}

func NewFileHash() *fileHash {
	return &fileHash{
		file: "",
		hash: &HashEncryption{},
	}
}

func (h *HashEncryption) FileMethod() FileMethod {
	return &fileHash{
		hash: &HashEncryption{
		},
	}
}

func (f *fileHash) FileMd5(file string, t string, wg *sync.WaitGroup) (string, error) {
	defer wg.Done()
	fileInfo, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	if fileInfo.IsDir() {
		return "", fmt.Errorf("file is directory")
	}

	fileObj, _ := os.Open(file)
	defer fileObj.Close()
	f.hash.wg.Add(1)
	go func() {
		f.file = f.hash.method(t, file, nil)
	}()
	f.hash.wg.Wait()
	return fmt.Sprintf("文件路径: %s, hash: %s", file, f.file), nil
}
