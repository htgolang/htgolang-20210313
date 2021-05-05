package line

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Line(file string) int {
	lines := 0
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		lines++
	}
	return lines
}

func DirLine(dir string) int {
	ret := 0
	fslist, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fs := range fslist {
		if fs.IsDir() {
			ret += DirLine(filepath.Join(dir, fs.Name()))
		} else {
			ret += Line(filepath.Join(dir, fs.Name()))
		}
	}
	return ret
}
