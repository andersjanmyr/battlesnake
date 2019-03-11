package main

import (
	"fmt"
	"net/http"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/gorilla/mux"
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

func Start(w http.ResponseWriter, r *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeRequest(r, &decoded)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dump(decoded)
	vars := mux.Vars(r)
	fmt.Println(vars)
	battleSnake := getBattleSnake(vars["kind"], vars["id"])
	respond(w, battleSnake.Start(&decoded))
}

func Move(w http.ResponseWriter, r *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeRequest(r, &decoded)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dump(decoded)

	vars := mux.Vars(r)
	battleSnake := getBattleSnake(vars["kind"], vars["id"])
	move := battleSnake.Move(&decoded)
	record(&decoded, move)
	respond(w, move)
}

func End(w http.ResponseWriter, r *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(r, &decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dump(decoded)
	record(&decoded, nil)
	respond(w, "This is the end beautiful friend")
}

func Ping(res http.ResponseWriter, req *http.Request) {
	respond(res, "pong")
}
