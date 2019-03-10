package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andersjanmyr/battlesnake/pkg/randy"
)

func main() {
	handleRoute("/", Index)
	handleRoute("/start", Start)
	handleRoute("/move", Move)
	handleRoute("/end", End)
	handleRoute("/ping", Ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	battlesnake = randy.New()

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	url := fmt.Sprintf("http://%s:%s/", IP(), port)
	log.Printf("Server started at url\n%s\n", url)
	http.ListenAndServe(":"+port, LoggingHandler(http.DefaultServeMux))
}

func handleRoute(path string, f func(w http.ResponseWriter, r *http.Request)) {
	http.Handle(path, LocalhostToIP(http.HandlerFunc(f)))
}
