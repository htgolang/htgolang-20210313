package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("table.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	line, err := reader.Read()
	fmt.Println(err, line)

	line, err = reader.Read()
	fmt.Println(err, line)

	line, err = reader.Read()
	fmt.Println(err, line)

	line, err = reader.Read()
	fmt.Println(err, line)

	file, _ = os.Create("table.csv")
	writer := csv.NewWriter(file)
	writer.Write([]string{"A1", "B1", "C1"})
	writer.Write([]string{"A2", "B2", "C2"})
	writer.Write([]string{"A3", "B3", "C3"})
	writer.Write([]string{"A4", "B4", "C4"})
	writer.Flush()
}
