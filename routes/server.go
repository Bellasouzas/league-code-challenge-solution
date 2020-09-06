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
	println("req:", req)
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
//func HandleInvert(res http.ResponseWriter, req *http.Request) {
//	file, _, err := req.FormFile("file")
//	if err != nil {
//		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
//		return
//	}
//	defer file.Close()
//
//	records, err := csv.NewReader(file).ReadAll()
//	if err != nil {
//		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
//		return
//	}
//
//	response := controllers.Invert(records)
//
//	fmt.Fprint(res, response, "\n")
//	return
//}

//HandleFlatten
//func HandleFlatten(res http.ResponseWriter, req *http.Request) {
//
//}

//HandleSum
//func HandleSum(res http.ResponseWriter, req *http.Request) {
//
//}

//HandleMultiply
//func HandleMultiply(res http.ResponseWriter, req *http.Request) {
//
//}
