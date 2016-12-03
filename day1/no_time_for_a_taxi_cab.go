package main

import "github.com/klauern/adventcode-2016/helpers"
import "strings"

// Position represents a position on a grid, all relative on an X/Y plane from point 0,0
type Position struct {
	X,
	Y int
	DirectionFacing string
}

func main() {
	input := helpers.MustLoadFile("input.txt")
	inputs := strings.Split(input, ", ")

	current := &Position{}
	for _, next := range inputs {
		current.CalculatePosition(next)
	}
}

func (p Position) CalculatePosition(next string) Position {
	return p
}
