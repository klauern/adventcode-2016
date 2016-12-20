package main

import (
	"strings"

	"strconv"

	"github.com/klauern/adventcode-2016/helpers"
)

type DestType int

const (
	bot DestType = iota
	output
	neither
)

type BotState struct {
	botMap      map[int][]int
	outputMap   map[int][]int
	movementMap map[int]*Movement
}

type Movement struct {
	highDest int
	lowDest  int
	highType DestType
	lowType  DestType
}

func (b *BotState) parseLine(line string) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "bot":
	case "value":

	}
}

func (b *BotState) addValue(line string) {
	fields := strings.Fields(line)
	value, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	botDest, err := strconv.Atoi(fields[5])
	if err != nil {
		panic(err)
	}
	b.botMap[botDest] = append(b.botMap[botDest], value)
}

func (b *BotState) addMovement(line string) {
	fields := strings.Fields(line)
	botNum, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	lowNum, err := strconv.Atoi(fields[6])
	if err != nil {
		panic(err)
	}
	lowDest := parseDest(fields[5])
	highNum, err := strconv.Atoi(fields[11])
	if err != nil {
		panic(err)
	}
	highDest := parseDest(fields[11])
	movement := &Movement{
		highDest: highNum,
		highType: highDest,
		lowDest:  lowNum,
		lowType:  lowDest,
	}
	b.movementMap[botNum] = movement
}

func parseDest(val string) DestType {
	switch val {
	case "output":
		return output
	case "bot":
		return bot
	}
	return neither
}

func main() {
	file := helpers.MustScanFile("input.txt")
	for file.Scan() {

	}
}
