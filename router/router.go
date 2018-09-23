package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

//Start starts a new router and resolves handlers
func Start(listenAddress string) {
	rtr := mux.NewRouter()

	addAppRoutes(rtr)

	go func() {
		err := http.ListenAndServe(listenAddress, rtr)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Listening on port:", listenAddress)
	}()
}
