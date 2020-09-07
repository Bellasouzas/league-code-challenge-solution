package matrix

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

//OpenCsv open and read csv file
func ReadCsv(csvFile io.Reader) [][]string {
	var matrix [][]string
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

func MatrixInt(records []string) []int {
	var matrixInt []int

	for _, val := range records {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return nil
		}
		matrixInt = append(matrixInt, valInt)
	}

	return matrixInt
}
