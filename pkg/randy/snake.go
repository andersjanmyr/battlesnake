package randy

import (
	"math/rand"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/andersjanmyr/battlesnake/pkg/core"
)

type snake struct {
	lastMove api.Move
}

var _ api.BattleSnake = &snake{}

// New creates a battlesnake
func New() api.BattleSnake {
	s := snake{}
	return &s
}

func (s *snake) Start(r *api.SnakeRequest) *api.StartResponse {
	return &api.StartResponse{
		HeadType: api.HeadFang,
		TailType: api.TailBlockBum,
	}
}

func (s *snake) Move(r *api.SnakeRequest) *api.MoveResponse {
	moves := core.PossibleMoves(r, s.lastMove)
	if len(moves) == 0 {
		return &api.MoveResponse{Move: s.lastMove}
	}
	random := rand.Intn(len(moves))
	s.lastMove = moves[random]
	return &api.MoveResponse{Move: s.lastMove}
}
