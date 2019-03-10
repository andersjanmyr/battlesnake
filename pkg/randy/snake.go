package randy

import (
	"fmt"
	"math/rand"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/andersjanmyr/battlesnake/pkg/core"
)

type snake struct {
	lastMove string
}

var _ api.BattleSnake = &snake{}

// New creates a battlesnake
func New() api.BattleSnake {
	s := snake{}
	return &s
}

func (s *snake) Start(r *api.StartRequest) *api.StartResponse {
	return &api.StartResponse{
		Color: "FF0000",
	}
}
func (s *snake) Move(r *api.Game) *api.MoveResponse {
	moves := core.PossibleMoves(r, s.lastMove)
	random := rand.Intn(len(moves))
	s.lastMove = moves[random]
	fmt.Println(moves, s.lastMove)
	return &api.MoveResponse{Move: s.lastMove}
}

func (s *snake) End(r *api.Game) string {
	return ""
}
