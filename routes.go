package main

import "net/http"

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	// Serve everything under /static/ directly from the static folder
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
}
