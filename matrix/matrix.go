package matrix

import (
	"encoding/csv"
	"io"
	"log"
)

//OpenCsv open and read csv file
func ReadCsv(csvFile io.Reader) [][]string {
	var matrix [][]string
	//read file
	records := csv.NewReader(csvFile)
	println(records)
	for {
		row, err := records.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

////CsvReaderRow parse row csv file
//func CsvReaderRow() {
//
//
//}
