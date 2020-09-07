package main

import (
	"fmt"
	"github.com/bellasouzas/league-code-challenge-solution/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//curl -F 'file=@./matrix/matrix.csv' "localhost:8080/echo"
//curl -F 'file=@./matrix/matrix.csv' "localhost:8080/invert"

const port = ":8000"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/echo", routes.HandleEcho)
	router.HandleFunc("/invert", routes.HandleInvert)
	router.HandleFunc("/flatten", routes.HandleFlatten)
	router.HandleFunc("/sum", routes.HandleSum)
	router.HandleFunc("/multiply", routes.HandleMultiply)

	fmt.Printf("Server is running at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
