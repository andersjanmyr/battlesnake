package horry

import (
	"fmt"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/andersjanmyr/battlesnake/pkg/core"
)

type snake struct {
	lastMove       api.Move
	secondLastMove api.Move
}

var _ api.BattleSnake = &snake{}

var horiz = []api.Move{api.Left, api.Right}

// New creates a battlesnake
func New() api.BattleSnake {
	s := snake{}
	return &s
}

func (s *snake) Start(r *api.SnakeRequest) *api.StartResponse {
	return &api.StartResponse{
		HeadType: api.HeadSmile,
		TailType: api.TailHook,
	}
}

func (s *snake) Move(r *api.SnakeRequest) *api.MoveResponse {
	moves := core.PossibleMoves(r, s.lastMove)
	if len(moves) == 0 {
		return &api.MoveResponse{Move: s.lastMove}
	}
	head := r.You.Body[0]
	if len(moves) > 1 && head.X < 2 {
		moves = core.Remove(moves, api.Left)
	}
	if len(moves) > 1 && head.X > r.Board.Width-3 {
		moves = core.Remove(moves, api.Right)
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

func (s *snake) moveResponse(move api.Move) *api.MoveResponse {
	if move != s.lastMove {
		s.secondLastMove = s.lastMove
	}
	s.lastMove = move
	fmt.Println(s)
	return &api.MoveResponse{Move: move}
}

func contains(moves []api.Move, move api.Move) bool {
	for _, m := range moves {
		if m == move {
			return true
		}
	}
	return false
}
