package routers

import "net/http"

//HandleEcho takes a csv file as request and returns a print in the console
// of the plain content of the file, line by line and comma separated
func HandleEcho(res http.ResponseWriter, req *http.Request) {

}

//HandleInvert takes a csv file as request and returns a print in the console
// of the same data but transposed, in other words, lines and columns
func HandleInvert(res http.ResponseWriter, req *http.Request) {

}

//HandleFlatten
func HandleFlatten(res http.ResponseWriter, req *http.Request) {

}

//HandleSum
func HandleSum(res http.ResponseWriter, req *http.Request) {

}

//HandleMultiply
func HandleMultiply(res http.ResponseWriter, req *http.Request) {

}
