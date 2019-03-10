package horry

import "github.com/andersjanmyr/battlesnake/api"

type snake struct{}

var _ api.BattleSnake = &snake{}

// New creates a battlesnake
func New() api.BattleSnake {
	s := snake{}
	return &s
}

func (s *snake) Start(r *api.StartRequest) *api.StartResponse {
	return &api.StartResponse{}
}
func (s *snake) Move(r *api.SnakeRequest) *api.MoveResponse {
	return &api.MoveResponse{Move: "down"}
}
