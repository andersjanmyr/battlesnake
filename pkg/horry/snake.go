package horry

import (
	"fmt"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/andersjanmyr/battlesnake/pkg/core"
)

type snake struct {
	lastMove       string
	secondLastMove string
}

var _ api.BattleSnake = &snake{}

var horiz = []string{"left", "right"}

// New creates a battlesnake
func New() api.BattleSnake {
	s := snake{}
	return &s
}

func (s *snake) Start(r *api.SnakeRequest) *api.StartResponse {
	return &api.StartResponse{
		HeadType: "smile",
		TailType: "hook",
	}
}

func (s *snake) Move(r *api.SnakeRequest) *api.MoveResponse {
	moves := core.PossibleMoves(r, s.lastMove)
	if len(moves) == 0 {
		return &api.MoveResponse{Move: s.lastMove}
	}
	if contains(horiz, s.lastMove) && contains(moves, s.lastMove) {
		return s.moveResponse(s.lastMove)
	}
	if contains(moves, "left") {
		return s.moveResponse("left")
	}
	if contains(moves, "right") {
		return s.moveResponse("right")
	}

	if contains(moves, s.secondLastMove) {
		return s.moveResponse(s.secondLastMove)
	}

	return s.moveResponse(moves[0])
}

func (s *snake) moveResponse(move string) *api.MoveResponse {
	if move != s.lastMove {
		s.secondLastMove = s.lastMove
	}
	s.lastMove = move
	fmt.Println(s)
	return &api.MoveResponse{Move: move}
}

func contains(moves []string, move string) bool {
	for _, m := range moves {
		if m == move {
			return true
		}
	}
	return false
}
