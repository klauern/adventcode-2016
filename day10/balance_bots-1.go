package main

import (
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

type BotState struct {
	botMap      map[int][]int
	outputMap   map[int][]int
	movementMap map[int]int
}

func (b *BotState) parseLine(line string) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "bot":
	case "value":

	}
}

func (b *BotState) addValue(line string) {

}

func (b *BotState) addMovement(line string) {

}

func main() {
	file := helpers.MustScanFile("input.txt")
	for file.Scan() {

	}
}
