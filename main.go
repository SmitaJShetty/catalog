package main

import (
	"AP/router"
	"fmt"
)

func main() {

	listenAddress := "localhost:8090"
	router.Start(listenAddress)
	fmt.Printf("Server listening on: %s ...", listenAddress)

	select {}
}
