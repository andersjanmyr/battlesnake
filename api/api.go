package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type SnakeRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type Game struct {
	ID string `json:"id"`
}

type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Coord `json:"food"`
	Snakes []Snake `json:"snakes"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Coord `json:"body"`
}

type StartResponse struct {
	Color    string `json:"color,omitempty"`
	HeadType string `json:"headType,omitempty"`
	TailType string `json:"tailType,omitempty"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

func DecodeSnakeRequest(req *http.Request, decoded *SnakeRequest) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
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
	Start(r *SnakeRequest) *StartResponse
	Move(r *SnakeRequest) *MoveResponse
}
