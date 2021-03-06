package main

import (
	"fmt"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

type KeyPos struct {
	X,
	Y int
}

func (k KeyPos) String() string {
	return fmt.Sprintf("%v %v", k.X, k.Y)

}

var deadSpot = KeyPos{X: -1, Y: -1}
var positionTree = [][]rune{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var crazyPositionTree = [][]rune{
	{'Z', 'Z', '1', 'Z', 'Z'},
	{'Z', '2', '3', '4', 'Z'},
	{'5', '6', '7', '8', '9'},
	{'Z', 'A', 'B', 'C', 'Z'},
	{'Z', 'Z', 'D', 'Z', 'Z'},
}

func (k *KeyPos) CalcMoves(move rune) KeyPos {
	switch move {
	case 'U':
		if k.Y == 0 {
			return KeyPos{k.X, k.Y}
		}
		return KeyPos{k.X, k.Y - 1}
	case 'D':
		if k.Y == 2 {
			return KeyPos{k.X, k.Y}
		}
		return KeyPos{k.X, k.Y + 1}
	case 'R':
		if k.X == 2 {
			return KeyPos{k.X, k.Y}
		}
		return KeyPos{k.X + 1, k.Y}
	case 'L':
		if k.X == 0 {
			return KeyPos{k.X, k.Y}
		}
		return KeyPos{k.X - 1, k.Y}
	}
	return deadSpot
}

func (k *KeyPos) CalcNextDigit(moves string) KeyPos {
	next := KeyPos{-1, -1}
	for _, move := range moves {
		if next == deadSpot {
			next = k.CalcMoves(move)
		} else {
			next = next.CalcMoves(move)
		}
		if next == deadSpot {
			panic("Not able to set")
		}
	}
	return next
}

func (k *KeyPos) GetDigit() int {
	return positionTree[k.Y][k.X]
}

func main() {
	input := helpers.MustLoadFile("input.txt")
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	key := &KeyPos{1, 1}
	digits := make([]int, 0)
	for _, digit := range lines {
		d := key.CalcNextDigit(digit)
		fmt.Printf("Digit is %v\n", d)
		digits = append(digits, d.GetDigit())
	}
	fmt.Printf("Code is %v", digits)
}
