package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var records [][]string

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err.Error())
		os.Exit(0)
	}
}

func readData(filePath string) {
	file, err := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err = csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err)
	defer file.Close()
}

func writeData(filePath string) {
	file, err := os.Create(filePath)
	checkError("Cannot create file "+filePath, err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	checkError("Cannot write to file "+filePath, err)
	defer file.Close()
}
