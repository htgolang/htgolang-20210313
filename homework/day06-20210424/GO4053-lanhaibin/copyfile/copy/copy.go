package copy

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Copy(src, dest string, block int) {
	fs, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()
	fd, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	reader := bufio.NewReaderSize(fs, block)
	writer := bufio.NewWriterSize(fd, block)
	defer writer.Flush()

	data := make([]byte, block)
	for {
		n, err := reader.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		writer.Write(data[:n])
	}
}
