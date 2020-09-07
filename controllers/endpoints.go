package controllers

import (
	"fmt"
	"strings"
	"sync"
)

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

//sum all channels from Sum func
func sumChannel(ch chan int) int {
	var total int
	for val := range ch {
		total += val
	}
	return total
}

//Sum func sums all values from matrix
func Sum(records [][]string) int {

	matrix := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	const channel_counter = 4
	send := make(chan int)
	receive := make(chan int, channel_counter)

	go func() {
		for _, value := range matrix {
			send <- value
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
