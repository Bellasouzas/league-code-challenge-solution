package league_code_challenge_solution

import (
	"github.com/bellasouzas/league-code-challenge-solution/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const address = ":8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/echo", routers.HandleEcho)
	//	router.HandleFunc("/echo", routers.HandleInvert)
	//	router.HandleFunc("/echo", routers.HandleFlatten)
	//	router.HandleFunc("/echo", routers.HandleSum)
	//	router.HandleFunc("/echo", routers.HandleMultiply)

	log.Fatal(http.ListenAndServe(address, nil))
}
