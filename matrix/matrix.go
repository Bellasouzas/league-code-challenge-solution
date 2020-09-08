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

//MatrixInt convert matrix [][]string to [][]int
func MatrixInt(matrix [][]string) [][]int {
	var rowItems [][]int
	for i := 0; i < len(matrix); i++ {
		var row []int
		for j := 0; j < len(matrix[i]); j++ {
			valInt, _ := strconv.Atoi(matrix[i][j])
			row = append(row, valInt)
		}
		rowItems = append(rowItems, row)
	}
	return rowItems
}

// MatrixSquare return if matrix is square
func MatrixIsSquare(matrix [][]string) bool {
	numRows := len(matrix)
	for _, row := range matrix {
		if len(row) == numRows {
			return true
		}
	}
	return false
}

//matrixIsEmpty return true if matrix is empty
func MatrixIsNotEmpty(matrix [][]string) bool {
	if len(matrix) == 0 {
		return false
	}
	return true
}
