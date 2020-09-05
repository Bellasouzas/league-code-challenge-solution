package cmd

import (
	"github.com/bellasouzas/league-code-challenge-solution/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/echo", routers.HandleEcho)
	router.HandleFunc("/echo", routers.HandleInvert)
	router.HandleFunc("/echo", routers.HandleFlatten)
	router.HandleFunc("/echo", routers.HandleSum)
	router.HandleFunc("/echo", routers.HandleMultiply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
