package main

import (
	"ArticleApp/router"
	"fmt"
)

func main() {

	listenAddress := "0.0.0.0:8090"
	router.Start(listenAddress)
	fmt.Printf("Server listening on: %s ...", listenAddress)

	select {}
}
