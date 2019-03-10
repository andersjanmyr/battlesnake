package core

import (
	"fmt"

	"github.com/andersjanmyr/battlesnake/api"
)

type Value int

const (
	Empty Value = iota
	Snake
	Food
	Wall
)

func PossibleMoves(req *api.SnakeRequest, direction api.Move) []api.Move {
	ms := remove(moves(), Opposite(direction))
	mz := []api.Move{}
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

func moves() []api.Move {
	return append(api.Moves[:0:0], api.Moves...)
}

func remove(l []api.Move, item api.Move) []api.Move {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

var opposites = map[api.Move]api.Move{
	api.Up:    api.Down,
	api.Down:  api.Up,
	api.Right: api.Left,
	api.Left:  api.Right,
}

func Opposite(dir api.Move) api.Move {
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

func nextCoord(coord api.Coord, direction api.Move) api.Coord {
	c := coord
	if direction == api.Up {
		c.Y = c.Y - 1
	}
	if direction == api.Down {
		c.Y = c.Y + 1
	}
	if direction == api.Right {
		c.X = c.X + 1
	}
	if direction == api.Left {
		c.X = c.X - 1
	}
	return c
}
