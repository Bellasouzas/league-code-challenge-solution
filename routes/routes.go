package routes

import (
	"fmt"
	"github.com/bellasouzas/league-code-challenge-solution/controllers"
	"github.com/bellasouzas/league-code-challenge-solution/matrix"
	"net/http"
)

func HandleEcho(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records := matrix.ReadCsv(file)

	notValidated := []bool{matrix.MatrixIsSquare(records), matrix.MatrixIsNotEmpty(records)}
	for i := 0; i < len(notValidated); i++ {
		if !notValidated[i] {
			fmt.Fprint(res, ("This file is not validated"))
			return
		}
	}

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

	notValidated := []bool{matrix.MatrixIsSquare(records), matrix.MatrixIsNotEmpty(records)}
	for i := 0; i < len(notValidated); i++ {
		if !notValidated[i] {
			fmt.Fprint(res, ("This file is not validated"))
			return
		}
	}

	response := controllers.Invert(records)

	fmt.Fprint(res, response, "\n")
	return
}

func HandleFlatten(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records := matrix.ReadCsv(file)

	notValidated := []bool{matrix.MatrixIsSquare(records), matrix.MatrixIsNotEmpty(records)}
	for i := 0; i < len(notValidated); i++ {
		if !notValidated[i] {
			fmt.Fprint(res, ("This file is not validated"))
			return
		}
	}

	response := controllers.Flatten(records)

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

	notValidated := []bool{matrix.MatrixIsSquare(records), matrix.MatrixIsNotEmpty(records)}
	for i := 0; i < len(notValidated); i++ {
		if !notValidated[i] {
			fmt.Fprint(res, ("This file is not validated"))
			return
		}
	}

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

	notValidated := []bool{matrix.MatrixIsSquare(records), matrix.MatrixIsNotEmpty(records)}
	for i := 0; i < len(notValidated); i++ {
		if !notValidated[i] {
			fmt.Fprint(res, ("This file is not validated"))
			return
		}
	}

	response := controllers.Multiply(records)

	fmt.Fprint(res, response, "\n")
	return
}
