package main

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/klauern/adventcode-2016/helpers"
)

// Position represents a position on a grid, all relative on an X/Y plane from point 0,0
type Position struct {
	X,
	Y int
	DirectionFacing rune
}

var allPositions map[Position]bool
var firstDouble Position

// Stringer prints the Position in a readable format.
func (p Position) String() string {
	return fmt.Sprintf("X: %v, Y: %v, Dir: %c", p.X, p.Y, p.DirectionFacing)
}

func main() {
	input := helpers.MustLoadFile("input.txt")
	input = strings.Trim(input, "\n ")

	pos := CalculateMovement(input)
	fmt.Printf("Position: %v\n", pos)

	fmt.Printf("Distance: %v\n", pos.CalculateDistanceFromHome())

	fmt.Printf("First Double Position: %v\n", firstDouble)
	fmt.Printf("Distance from Double Position: %v\n", firstDouble.CalculateDistanceFromHome())
}

// CalculateMovement will calculate the final position of a piece.
func CalculateMovement(moves string) *Position {
	inputs := strings.Split(moves, ", ")

	current := &Position{0, 0, 'N'}
	allPositions = make(map[Position]bool)
	allPositions[*current] = true
	for _, next := range inputs {
		current = current.CalculatePosition(next)
		if allPositions[*current] {
			firstDouble = *current
		} else {
			allPositions[*current] = true
		}
	}

	return current
}

// CalculatePosition will calculate the position of the piece after the move given in the 'next' string.
func (p *Position) CalculatePosition(next string) *Position {
	dir := p.CalculateDirection(next)
	distance, err := strconv.Atoi(next[1:])
	if err != nil {
		panic(err)
	}
	x := p.X
	y := p.Y
	switch dir {
	case 'N':
		y += distance
	case 'E':
		x += distance
	case 'W':
		x -= distance
	case 'S':
		y -= distance
	}
	return &Position{x, y, dir}
}

// CalculateDirection will calculate the next Position of an entry
// given the movement string.
func (p *Position) CalculateDirection(next string) rune {
	turn := next[0]
	fmt.Printf("Turning %c\n", turn)
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
	panic(fmt.Errorf("Unknown Direction: %c", turn))
}

// CalculateDistanceFromHome will calculate the overall distance from the starting point using
// cartesian coordinate calculations.
func (p *Position) CalculateDistanceFromHome() int {
	var x, y int
	if p.X < 0 {
		x = p.X * -1
	} else {
		x = p.X
	}
	if p.Y < 0 {
		y = p.Y * -1
	} else {
		y = p.Y
	}
	return x + y
}
