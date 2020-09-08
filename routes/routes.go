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
