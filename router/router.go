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

		rtr.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			fmt.Println(t)
			return nil
		})

		fmt.Println("Listening on port:", listenAddress)
	}()
}
