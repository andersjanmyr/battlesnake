package main

import (
	"log"
	"net/http"

	"github.com/andersjanmyr/battlesnake/api"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`
	<title>Anders Janmyr's Battlesnake</title>
	<h1>Anders Janmyr Battlesnake</h1>
	<a href="https://github.com/andersjanmyr/battlesnake">https://github.com/andersjanmyr/battlesnake</a>
	<h2>Routes</h2>
	<ul>
	<li><a href="/start">/start</a></li>
	<li><a href="/move">/move</a></li>
	<li><a href="/end">/end</a></li>
	<li><a href="/ping">/ping</a></li>
	</ul>
	`))
}

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	dump(decoded)

	respond(res, api.MoveResponse{
		Move: "down",
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	dump(decoded)
	respond(res, "That's all folks!")
}

func Ping(res http.ResponseWriter, req *http.Request) {
	respond(res, "pong")
}
