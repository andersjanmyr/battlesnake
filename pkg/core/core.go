package core

import (
	"fmt"

	"github.com/andersjanmyr/battlesnake/api"
)

var Moves = []string{"up", "down", "left", "right"}

var Heads = []string{
	"regular",
	"beluga",
	"bendr",
	"dead",
	"evil",
	"fang",
	"pixel",
	"safe",
	"sand-worm",
	"shades",
	"silly",
	"smile",
	"tongue",
}

var Tails = []string{
	"regular",
	"block-bum",
	"bolt",
	"curled",
	"fat-rattle",
	"freckled",
	"hook",
	"pixel",
	"round-bum",
	"sharp",
	"skinny",
	"small-rattle",
}

type Value int

const (
	Empty Value = iota
	Snake
	Food
	Wall
)

func PossibleMoves(req *api.SnakeRequest, direction string) []string {
	ms := remove(moves(), Opposite(direction))
	mz := []string{}
	for _, m := range ms {
		c := nextCoord(req.You.Body[0], m)
		fmt.Println("next", req.You.Body[0], m, c)
		v := ValueAt(req.Board, c)
		if v == Food || v == Empty {
			mz = append(mz, m)
		}
	}
	return mz
}

func moves() []string {
	return append(Moves[:0:0], Moves...)
}

func remove(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

var opposites = map[string]string{
	"up":    "down",
	"down":  "up",
	"left":  "right",
	"right": "left",
}

func Opposite(dir string) string {
	return opposites[dir]
}

func ValueAt(b api.Board, c api.Coord) Value {
	if c.X < 0 || c.Y < 0 || c.X >= b.Width || c.Y >= b.Height {
		return Wall
	}
	if coordTaken(b.Food, c) {
		return Food
	}
	for _, s := range b.Snakes {
		if coordTaken(s.Body, c) {
			return Snake
		}
	}
	return Empty
}

func coordTaken(cs []api.Coord, v api.Coord) bool {
	for _, c := range cs {
		if c == v {
			return true
		}
	}
	return false
}

func nextCoord(coord api.Coord, direction string) api.Coord {
	c := coord
	if direction == "up" {
		c.Y = c.Y - 1
	}
	if direction == "down" {
		c.Y = c.Y + 1
	}
	if direction == "right" {
		c.X = c.X + 1
	}
	if direction == "left" {
		c.X = c.X - 1
	}
	return c
}
