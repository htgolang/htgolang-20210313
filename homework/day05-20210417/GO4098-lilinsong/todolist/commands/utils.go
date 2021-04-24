package commands

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const fileSize = 1024

type myFileInfoList []os.FileInfo

func (m myFileInfoList) Len() int {
	return len(m)
}

func (m myFileInfoList) Less(i, j int) bool {
	return m[j].ModTime().Sub(m[i].ModTime()) > 0
}

func (m myFileInfoList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func getSortedFile() myFileInfoList {
	currentPath, _ := os.Getwd()
	var fileInfos myFileInfoList
	path, _ := os.Open(currentPath + "/data")
	files, _ := path.Readdir(-1)

	if len(files) == 0 {
		return []os.FileInfo{}
	}

	for _, file := range files {
		fileInfos = append(fileInfos, file)
	}
	sort.Sort(fileInfos)
	return fileInfos
}

func dataDir() {
	currentPath, _ := os.Getwd()
	stat, err := os.Stat(currentPath + "/data")
	if os.IsNotExist(err) {
		os.Mkdir(currentPath+"/data", 0755)
	} else if !stat.IsDir() {
		fmt.Println("存在名为data文件，无法创建保存的目录")
		return
	}
}

func Load() {
	currentPath, _ := os.Getwd()
	dataDir()
	fileInfos := getSortedFile()
	file := fileInfos[len(fileInfos)-1]
	fileName := filepath.Join(currentPath+"/data", file.Name())
	JsonLoad(fileName)
}

func BakFile() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		fileLayout := "20060102"
		timeStamp := fmt.Sprintf("tasks_%s.json", time.Now().Format(fileLayout))
		path, _ := os.Getwd()
		newFile, _ := os.Create(filepath.Join(path+"/data", timeStamp))
		defer newFile.Close()

		oldFile, _ := ioutil.ReadFile(filepath.Join(path, "data/tasks.json"))

		buffer := make([]byte, 512)
		io.CopyBuffer(newFile, strings.NewReader(string(oldFile)), buffer)
	})
	c.Start()
}

/*
func FileSize(file *os.File) {
	var fileLayout = "20060102"
	fileInfo, _ := file.Stat()
	if fileInfo.Size() < fileSize {
		return
	}

	now := time.Now()
	timeStamp := fmt.Sprintf("%s_%d", now.Format(fileLayout), now.Unix())
	path, _ := os.Getwd()
	timeFile := filepath.Join(path+"/data", timeStamp+".json")
	newFile, _ := os.OpenFile(timeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	buffer := make([]byte, 512)
	if _, err := io.CopyBuffer(newFile, file, buffer); err != nil {
		fmt.Println(err)
	}

	defer newFile.Close()

	oldFile := filepath.Join(path, "data/tasks.json")
	os.Truncate(oldFile, 0)
	file.Seek(0, 0)

}
*/
