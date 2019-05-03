package hungry

import (
	"math"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/andersjanmyr/battlesnake/pkg/core"
)

type snake struct {
	lastMove api.Move
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
		HeadType: api.HeadBendr,
		TailType: api.TailFreckled,
	}
}

func (s *snake) Move(r *api.SnakeRequest) *api.MoveResponse {
	moves := core.PossibleMoves(r, s.lastMove)
	if len(moves) == 0 {
		return s.moveResponse(s.lastMove)
	}
	if len(moves) == 1 {
		return s.moveResponse(moves[0])
	}
	head := r.You.Body[0]
	food := findClosestFood(head, r.Board.Food)
	if head.X < food.X && contains(moves, api.Right) {
		return s.moveResponse(api.Right)
	}
	if head.X > food.X && contains(moves, api.Left) {
		return s.moveResponse(api.Left)
	}
	if head.Y < food.Y && contains(moves, api.Down) {
		return s.moveResponse(api.Down)
	}
	if head.Y > food.Y && contains(moves, api.Up) {
		return s.moveResponse(api.Up)
	}
	return s.moveResponse(moves[0])
}

func (s *snake) moveResponse(move api.Move) *api.MoveResponse {
	s.lastMove = move
	return &api.MoveResponse{Move: move}
}

func findClosestFood(head api.Coord, food []api.Coord) api.Coord {
	closestDistance := math.MaxFloat64
	closestCoord := api.Coord{}
	for _, c := range food {
		if d := distance(head, c); d < closestDistance {
			closestCoord = c
			closestDistance = d
		}
	}
	return closestCoord
}

func distance(a, b api.Coord) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func contains(moves []api.Move, move api.Move) bool {
	for _, m := range moves {
		if m == move {
			return true
		}
	}
	return false
}
