package main

import (
	"testing"
)

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 2, Y: 0},
		Body: []Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
		}
	}
}

func TestBodyAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		Head: Coord{X: 8, Y: 1},
		Body: []Coord{{X: 7, Y: 1}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		You: me,
	}
	nextMove := move(state)
	// Assert never move left
	if nextMove.Move == "left" {
		t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
	}
	if nextMove.Move == "down" {
		t.Errorf("snake moved onto its own body, %s", nextMove.Move)
	}
}

func TestManhattan(t *testing.T) {
	a := Coord{X: 5, Y: 10}
	b := Coord{X: 1, Y: 5}

	if manHattanDistance(a, b) != 9 {
		t.Errorf("manhattan distance should equal 9, but was %d", manHattanDistance(a, b))
	}
}
