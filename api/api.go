package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Coords struct {
	Data []Coord
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snakes struct {
	Data []Snake
}

type Snake struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Health int    `json:"health"`
	Body   Coords `json:"body"`
}

type DeadSnakes struct {
	Data []DeadSnake
}

type DeadSnake struct {
	ID     string `json:"id"`
	Length string `json:"length"`
	Death  int    `json:"health"`
}

type Death struct {
	Ture   int      `json:"turn"`
	causes []string `json:"causes"`
}

type Game struct {
	ID     int    `json:"id"`
	Turn   int    `json:"turn"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Food   Coords `json:"food"`
	Snakes Snakes `json:"snakes"`
	You    Snake  `json:"you"`
}

type StartRequest struct {
	GameID int `json:"game_id"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type StartResponse struct {
	Color          string `json:"color,omitempty"`
	SecondaryColor string `json:"secondary_color,omitempty"`
	HeadURL        string `json:"head_url,omitempty"`
	Taunt          string `json:"taunt,omitempty"`
	HeadType       string `json:"head_type,omitempty"`
	TailType       string `json:"tail_type,omitempty"`
}

type EndRequest struct {
	GameID     int        `json:"game_id"`
	Winners    []string   `json:"winners"`
	DeadSnakes DeadSnakes `json:"dead_snakes"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

func NewCoords(c Coord) Coords {
	return Coords{Data: []Coord{c}}
}

func NewMoveResponse(move string) *MoveResponse {
	return &MoveResponse{Move: move}
}

func DecodeRequest(req *http.Request, decoded interface{}) error {
	requestDump, err2 := httputil.DumpRequest(req, true)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("DECODE", string(requestDump))
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}

type BattleSnake interface {
	Start(r *StartRequest) *StartResponse
	Move(r *Game) *MoveResponse
}
