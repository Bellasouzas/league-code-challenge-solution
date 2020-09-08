package controllers

import (
	"fmt"
	"github.com/bellasouzas/league-code-challenge-solution/matrix"
	"strings"
	"sync"
)

//todo refatorar : trocar o nome do arquivo para controllers.go

// Echo takes a csv file as request and returns a print in the console
// of the plain content of the file, line by line and comma separated
func Echo(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

//Invert receives matrix NxN and transposed, in other words, lines and columns
func Invert(records [][]string) string {
	var response string
	xl := len(records[0])
	yl := len(records)
	inverted := make([][]string, xl)

	for i := range inverted {
		inverted[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			inverted[i][j] = records[j][i]
		}

	}
	for _, row := range inverted {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}

func Flatten(records [][]string) string {
	var list string
	var response string

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			list = list + records[i][j] + ","
		}
	}
	response = list[0 : len(list)-1]
	return response
}

//sum all channels
func sumChannel(ch chan int) int {
	var total int
	for val := range ch {
		total += val
	}
	return total
}

//Sum func sums all values from matrix
func Sum(records [][]string) int {
	valuesToSum := matrix.MatrixInt(records)
	const channel_counter = 4
	send := make(chan int)
	receive := make(chan int, channel_counter)

	go func() {
		for i := 0; i < len(valuesToSum); i++ {
			for j := 0; j < len(valuesToSum[i]); j++ {
				send <- valuesToSum[i][j]
			}
		}
		close(send)
	}()
	var wg sync.WaitGroup
	wg.Add(channel_counter)

	for i := 0; i < channel_counter; i++ {
		go func() {
			receive <- sumChannel(send)
			wg.Done()
		}()
	}
	wg.Wait()
	close(receive)
	response := sumChannel(receive)
	return response
}

//Nultiply func sums all values from matrix
func Multiply(records [][]string) int {
	valuesToMultiply := matrix.MatrixInt(records)
	multipy_counter := 1

	for i := 0; i < len(valuesToMultiply); i++ {
		for j := 0; j < len(valuesToMultiply[i]); j++ {
			multipy_counter *= valuesToMultiply[i][j]
		}
	}
	response := multipy_counter
	return response
}
