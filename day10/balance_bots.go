package main

import (
	"fmt"
	"strconv"
	"strings"

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

func (b *BotState) addLineAssignment(line string) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "bot":
		b.addMovement(line)
	case "value":
		b.addValue(line)
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

func (b *BotState) calcBotMove() {
	for k, v := range b.botMap {
		if len(v) == 2 {
			b.move(k)
		}
	}
}

func (b *BotState) move(botID int) {
	movement := b.movementMap[botID]
	high, low := highLow(b.botMap[botID])
	if movement.lowType == output {
		b.outputMap[movement.lowDest] = append(b.outputMap[movement.lowDest], low)
	} else if movement.lowType == bot {
		b.botMap[movement.lowDest] = append(b.botMap[movement.lowDest], low)
	}
	if movement.highType == output {
		b.outputMap[movement.highDest] = append(b.outputMap[movement.highDest], high)
	} else if movement.highType == bot {
		b.botMap[movement.highDest] = append(b.botMap[movement.highDest], high)
	}
}

func highLow(vals []int) (int, int) {
	if vals[0] > vals[1] {
		return vals[0], vals[1]
	}
	return vals[1], vals[0]
}

func main() {
	botState := &BotState{botMap: make(map[int][]int, 0), movementMap: make(map[int]*Movement, 0), outputMap: make(map[int][]int, 0)}
	file := helpers.MustScanFile("input.txt")
	for file.Scan() {
		botState.addLineAssignment(file.Text())
	}
	for k, v := range botState.botMap {
		fmt.Printf("Bot %v has %v chips\n", k, v)
	}

}
