package main

import (
	"fmt"
	"net/http"
)

func main() {
	// init mux router
	mux := http.NewServeMux()
	registerRoutes(mux)
	httpServer := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	// Server listen
	if err := httpServer.ListenAndServe();
	err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Server is running on C")
}
