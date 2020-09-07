package routes

import (
	"fmt"
	"github.com/bellasouzas/league-code-challenge-solution/controllers"
	"github.com/bellasouzas/league-code-challenge-solution/matrix"
	"net/http"
)

//HandleEcho takes a csv file as request and returns a print in the console
// of the plain content of the file, line by line and comma separated
func HandleEcho(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records := matrix.ReadCsv(file)
	response := controllers.Echo(records)

	fmt.Fprint(res, response)

}

//HandleInvert takes a csv file as request and returns a print in the console
// of the same data but transposed, in other words, lines and columns
func HandleInvert(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records := matrix.ReadCsv(file)
	response := controllers.Invert(records)

	fmt.Fprint(res, response, "\n")
	return
}

//HandleFlatten
func HandleFlatten(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records := matrix.ReadCsv(file)
	response := controllers.Flatten(records)
	// Asserting that the input matrix is NxN dimentional
	var inputTest = 0
	for i := 0; i < len(records); i++ {
		if len(records[i]) != len(records) {
			inputTest++
		}
	}
	if inputTest != 0 {
		fmt.Fprint(res, "The input matrix needs to be NxN dimention\n")
		return
	}
	fmt.Fprint(res, response, "\n")
	return
}

//HandleSum sums all values from matrix
func HandleSum(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records := matrix.ReadCsv(file)
	response := controllers.Sum(records)

	fmt.Fprint(res, response, "\n")
	return
}

//HandleMultiply sends values from records multiplied
func HandleMultiply(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records := matrix.ReadCsv(file)
	response := controllers.Multiply(records)

	fmt.Fprint(res, response, "\n")
	return
}
