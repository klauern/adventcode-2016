package main

import (
	"github.com/klauern/adventcode-2016/helpers"
)

type KeyPos struct {
	X,
	Y int
}

var deadSpot = KeyPos{X: -1, Y: -1}
var positionTree = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

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
	return KeyPos{-1, -1}
}

func (k *KeyPos) CalcNextDigit(moves string) KeyPos {
	next := KeyPos{-1, -1}
	for _, move := range moves {
		if next == deadSpot {
			next = k.CalcMoves(move)
		} else {
			next = next.CalcMoves(move)
		}
	}
	return next
}

func (k *KeyPos) GetDigit() int {
	return positionTree[k.Y][k.X]
}

func main() {
	helpers.MustLoadFile("input.txt")

}
