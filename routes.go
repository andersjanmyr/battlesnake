package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andersjanmyr/battlesnake/api"
)

var battlesnake api.BattleSnake

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

	respond(w, battlesnake.Start(&decoded))
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

	respond(w, battlesnake.Move(&decoded))
}

func End(w http.ResponseWriter, r *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(r, &decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dump(decoded)
	respond(w, "This is the end beautiful friend")
}

func Ping(res http.ResponseWriter, req *http.Request) {
	respond(res, "pong")
}

func LocalhostToIP(next http.Handler) http.Handler {
	ip := IP()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, port := getHostPort(r.Host)
		url := fmt.Sprintf("http://%s%s%s", ip, port, r.URL.Path)
		if host == "127.0.0.1" || host == "localhost" {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getHostPort(s string) (string, string) {
	ss := strings.Split(s, ":")
	if len(ss) < 2 {
		return ss[0], ""
	}
	return ss[0], ":" + ss[1]
}
