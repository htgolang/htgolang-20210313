package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	err := follow("./test.log")
	fmt.Print(err)
}

func follow(filename string) error {
	file, _ := os.Open(filename)
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	_ = watcher.Add(filename)

	r := bufio.NewReader(file)
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Print(string(by))
		if err != io.EOF {
			continue
		}
		if err = waitForChange(watcher); err != nil {
			return err
		}
	}
}

func waitForChange(w *fsnotify.Watcher) error {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-w.Errors:
			return err
		}
	}
}
