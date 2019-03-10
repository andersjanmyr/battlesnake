package api

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockReq = `{
	"game": {
		"id": "abc"
	},
	"turn": 10,
	"board": {
		"height": 10,
		"width": 10,
		"food": [{
				"x": 1,
				"y": 1
		}],
		"snakes": [{
			"id": "123",
			"name": "snek",
			"health": 1,
			"body": [{
					"x": 1,
					"y": 1
			}]
		}]
	},
	"you": {
		"id": "123",
		"name": "snek",
		"health": 1,
		"body": [{
				"x": 1,
				"y": 1
		}]
	}
}`

func requestWithBody(body string) *http.Request {
	req, err := http.NewRequest("", "", bytes.NewBufferString(body))
	if err != nil {
		panic(err)
	}
	return req
}

func createMockGame() Game {
	c := Coord{
		X: 1,
		Y: 1,
	}
	s := Snake{
		ID:     "123",
		Name:   "snek",
		Health: 1,
		Body:   NewCoords(c),
	}
	sr := Game{
		ID:     666,
		Height: 10,
		Width:  10,
		Food:   NewCoords(c),
		Snakes: []Snake{s},
		Turn:   10,
		You:    s,
	}
	return sr
}

func TestDecodeSnakeRequest(t *testing.T) {
	expected := Game{}
	req := requestWithBody(mockReq)

	result := Game{}
	err := DecodeRequest(req, &result)
	if assert.NoError(t, err) {
		expected = createMockGame()
	}
	assert.Equal(t, result, expected)
}
