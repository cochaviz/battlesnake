package internal

import (
	"battlesnake/internal/utils"
	"battlesnake/pkg/api"
	"testing"
)

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := api.Battlesnake{
		// Length 3, facing right
		Head: api.Coord{X: 2, Y: 0},
		Body: []api.Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := api.GameState{
		Board: api.Board{
			Snakes: []api.Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove, _ := Move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
		}
	}
}

func TestBodyAvoidance(t *testing.T) {
	// Arrange
	me := api.Battlesnake{
		Head: api.Coord{X: 8, Y: 1},
		Body: []api.Coord{{X: 7, Y: 1}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0}},
	}
	state := api.GameState{
		Board: api.Board{
			Snakes: []api.Battlesnake{me},
		},
		You: me,
	}
	nextMove, _ := Move(state)
	// Assert never move left
	if nextMove.Move == "left" {
		t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
	}
	if nextMove.Move == "down" {
		t.Errorf("snake moved onto its own body, %s", nextMove.Move)
	}
}

func TestManhattan(t *testing.T) {
	a := api.Coord{X: 5, Y: 10}
	b := api.Coord{X: 1, Y: 5}

	if utils.ManHattanDistance(a, b) != 9 {
		t.Errorf("manhattan distance should equal 9, but was %d", utils.ManHattanDistance(a, b))
	}
}
