package api

import (
	"encoding/json"
	"net/http"
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
	HeadType Head   `json:"headType,omitempty"`
	TailType Tail   `json:"tailType,omitempty"`
}

type MoveResponse struct {
	Move Move `json:"move"`
}

type Move string

const (
	Up    Move = "up"
	Down  Move = "down"
	Left  Move = "left"
	Right Move = "right"
)

var Moves = []Move{Up, Down, Left, Right}

type Head string

const (
	HeadBeluga   Head = "beluga"
	HeadBendr    Head = "bendr"
	HeadDead     Head = "dead"
	HeadEvil     Head = "evil"
	HeadFang     Head = "fang"
	HeadPixel    Head = "pixel"
	HeadRegular  Head = "regular"
	HeadSafe     Head = "safe"
	HeadSandWorm Head = "sand-worm"
	HeadSilly    Head = "silly"
	HeadSmile    Head = "smile"
	HeadTongue   Head = "tongue"
)

var Heads = []Head{
	HeadBeluga,
	HeadBendr,
	HeadDead,
	HeadEvil,
	HeadFang,
	HeadPixel,
	HeadRegular,
	HeadSafe,
	HeadSandWorm,
	HeadSilly,
	HeadSmile,
	HeadTongue,
}

type Tail string

const (
	TailRegular     Tail = "regular"
	TailBlockBum    Tail = "block-bum"
	TailBolt        Tail = "bolt"
	TailCurled      Tail = "curled"
	TailFatRattle   Tail = "fat-rattle"
	TailFreckled    Tail = "freckled"
	TailHook        Tail = "hook"
	TailPixel       Tail = "pixel"
	TailRoundBum    Tail = "round-bum"
	TailSharp       Tail = "sharp"
	TailSkinny      Tail = "skinny"
	TailSmallRattle Tail = "small-rattle"
)

var Tails = []Tail{
	TailRegular,
	TailBlockBum,
	TailBolt,
	TailCurled,
	TailFatRattle,
	TailFreckled,
	TailHook,
	TailPixel,
	TailRoundBum,
	TailSharp,
	TailSkinny,
	TailSmallRattle,
}

func DecodeSnakeRequest(req *http.Request, decoded *SnakeRequest) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}

func DecodeRequest(req *http.Request, decoded interface{}) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}

type BattleSnake interface {
	Start(r *SnakeRequest) *StartResponse
	Move(r *SnakeRequest) *MoveResponse
}
