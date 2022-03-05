package data

import (
	"battlesnake/pkg/api"
	"log"
	"testing"
)

func TestTerminal(t *testing.T) {
	// Arrange
	me := api.Battlesnake{
		// Length 3, facing right
		Head: api.Coord{X: 0, Y: 0},
		Body: []api.Coord{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1}},
	}
	state := GameState{
		Board: api.Board{
			Snakes: []api.Battlesnake{me},
			Width:  2,
			Height: 2,
		},
		You: me,
	}
	state.Init()

	if !state.IsTerminal() {
		t.Error("State should be terminal")
	}
}

func TestMove(t *testing.T) {
	// Arrange
	me := api.Battlesnake{
		// Length 3, facing right
		Head: api.Coord{X: 1, Y: 0},
		Body: []api.Coord{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1}},
	}
	state := GameState{
		Board: api.Board{
			Snakes: []api.Battlesnake{me},
			Width:  2,
			Height: 2,
		},
		You: me,
	}
	state.Init()

	if state.IsTerminal() {
		t.Error("State should not be terminal")
	}
	safeMoves := state.SafeMoves()

	if len(safeMoves) != 1 {
		t.Error("There should be exactly 1 safe move")
	}
	if safeMoves[0] != "left" {
		t.Errorf("The safe move should be left, but was %s", safeMoves[0])
	}

	newState, err := state.Move(safeMoves[0])
	log.Print(state.You)
	log.Print(newState.You)

	if newState.You.Head.X != 0 || newState.You.Head.Y != 0 {
		t.Errorf("New head position should be (0,0), but was: (%d, %d)", newState.You.Head.X, newState.You.Head.Y)
	}

	if err != nil || newState.IsTerminal() {
		t.Error("State should not be terminal")
	}
	newSafeMoves := newState.SafeMoves()

	if newSafeMoves[0] != "up" {
		t.Errorf("The safe move should be right, but was %s", newSafeMoves[0])
	}
}

func TestMoveMore(t *testing.T) {
	// Arrange
	me := api.Battlesnake{
		// Length 3, facing right
		Head: api.Coord{X: 1, Y: 1},
		Body: []api.Coord{{X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := GameState{
		Board: api.Board{
			Snakes: []api.Battlesnake{me},
			Width:  3,
			Height: 3,
		},
		You: me,
	}
	state.Init()

	if state.IsTerminal() {
		t.Error("State should not be terminal")
	}
	safeMoves := state.SafeMoves()

	if len(safeMoves) != 3 {
		t.Error("There should be exactly 3 safe move")
	}
}

func TestMoveFood(t *testing.T) {
	// Arrange
	me := api.Battlesnake{
		// Length 3, facing right
		Head: api.Coord{X: 1, Y: 0},
		Body: []api.Coord{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1}},
	}
	state := GameState{
		Board: api.Board{
			Food:   []api.Coord{{X: 0, Y: 0}},
			Snakes: []api.Battlesnake{me},
			Width:  2,
			Height: 2,
		},
		You: me,
	}
	state.Init()

	if state.IsTerminal() {
		t.Error("State should not be terminal")
	}
	safeMoves := state.SafeMoves()

	if len(safeMoves) != 1 {
		t.Error("There should be exactly 1 safe move")
	}
	if safeMoves[0] != "left" {
		t.Errorf("The safe move should be left, but was: %s", safeMoves[0])
	}
	newState, err := state.Move(safeMoves[0])

	if err != nil || !newState.IsTerminal() {
		t.Errorf("State should be terminal, but there were possible moves: %s", newState.SafeMoves())
	}
}
