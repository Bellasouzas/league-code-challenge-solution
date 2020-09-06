package controllers

import (
	"fmt"
	"strings"
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
