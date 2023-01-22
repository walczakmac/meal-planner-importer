package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Csv struct {
	filepath string
}

func NewCsv(filepath string) *Csv {
	return &Csv{filepath: filepath}
}

func (c Csv) readFile() [][]string {
	file, err := os.Open(c.filepath)
	if nil != err {
		log.Fatal(err)
		return [][]string{}
	}

	reader := csv.NewReader(file)

	csvRecords, err := reader.ReadAll()
	if nil != err {
		log.Fatal(err)
		return [][]string{}
	}

	return csvRecords
}
