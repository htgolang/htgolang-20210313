package copy

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Copy(src, dest string) error {
	fs, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fs.Close()

	// 保留权限
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	fd, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fi.Mode())
	if err != nil {
		return err
	}
	defer fd.Close()
	// reader := bufio.NewReader(fs)
	writer := bufio.NewWriter(fd)
	defer writer.Flush()
	_, err = writer.ReadFrom(fs)
	return err
}

func CopyDir(src, dest string) error {
	fsi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !fsi.IsDir() {
		return fmt.Errorf("%s is Not a directory", fsi.Name())
	}
	fdi, err := os.Stat(dest)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = os.Mkdir(dest, fsi.Mode())
		if err != nil {
			return err
		}
	} else {
		if !fdi.IsDir() {
			return fmt.Errorf("%s is Not a directory", fdi.Name())
		}
	}

	fslist, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, fs := range fslist {
		if fs.IsDir() {
			CopyDir(filepath.Join(src, fs.Name()), filepath.Join(dest, fs.Name()))
		} else {
			Copy(filepath.Join(src, fs.Name()), filepath.Join(dest, fs.Name()))
		}
	}
	return nil
}
