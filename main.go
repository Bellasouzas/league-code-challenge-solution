package main

import (
	"fmt"
	"github.com/bellasouzas/league-code-challenge-solution/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//curl -F 'file=@./matrix/matrix.csv' "localhost:8080/echo"

const port = ":8080"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/echo", routes.HandleEcho)
	//router.HandleFunc("/invert", routes.HandleInvert)
	//	router.HandleFunc("/echo", routes.HandleFlatten) router.HandleFunc("/echo",
	//	routes.HandleSum) router.HandleFunc("/echo", routes.HandleMultiply)

	fmt.Printf("Server is running at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
