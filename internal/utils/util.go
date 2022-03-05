package utils

import (
	"battlesnake/pkg/api"
)

func IntAbs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func ManHattanDistance(a api.Coord, b api.Coord) int {
	return IntAbs(a.X-b.X) + IntAbs(a.Y-b.Y)
}
