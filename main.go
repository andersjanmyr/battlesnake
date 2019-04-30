package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(LoggingHandler, LocalhostToIP)
	router.HandleFunc("/", Index)
	addRoutes("/", router)
	addRoutes("/{kind}/{id}", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	url := fmt.Sprintf("http://%s:%s/{kind}/{id}", IP(), port)
	log.Printf("Server started at url\n%s\n", url)
	log.Printf("Available snakes are: %s\n", GetSnakeKinds())
	http.ListenAndServe(":"+port, router)
}

func addRoutes(path string, router *mux.Router) {
	var r *mux.Router
	if path == "/" {
		r = router
	} else {
		r = router.Methods("POST").PathPrefix(path).Subrouter()
	}
	r.HandleFunc("/start", Start)
	r.HandleFunc("/move", Move)
	r.HandleFunc("/end", End)
	r.HandleFunc("/ping", Ping)
}
