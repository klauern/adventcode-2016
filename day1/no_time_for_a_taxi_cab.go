package main

import (
	"fmt"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

// Position represents a position on a grid, all relative on an X/Y plane from point 0,0
type Position struct {
	X,
	Y int
	DirectionFacing rune
}

func main() {
	input := helpers.MustLoadFile("input.txt")
	inputs := strings.Split(input, ", ")

	current := &Position{}
	for _, next := range inputs {
		current.CalculateDirection(next)
	}
}

func (p Position) CalculatePosition(next string) Position {
	return p
}

// CalculateDirection will calculate the next Position of an entry
// given the movement string.
func (p Position) CalculateDirection(next string) rune {
	turn := next[0]
	fmt.Printf("Turning %x\n", turn)
	switch p.DirectionFacing {
	case 'N':
		if turn == 'L' {
			return 'W'
		}
		return 'E'
	case 'E':
		if turn == 'L' {
			return 'N'
		}
		return 'S'
	case 'W':
		if turn == 'L' {
			return 'S'
		}
		return 'N'
	case 'S':
		if turn == 'L' {
			return 'E'
		}
		return 'W'
	}
	return 'G'
}
