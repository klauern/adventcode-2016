package main

import (
	"strconv"
	"strings"
)

type value int
type bot int
type output int

type valueMove struct {
	value,
	bot int
}

type Destination interface {
	isOutput() bool
	isBot() bool
}

type movement struct {
	lowTo  Destination
	highTo Destination
}

func (b bot) isOutput() bool {
	return false
}

func (b bot) isBot() bool {
	return true
}

func (o output) isOutput() bool {
	return true
}

func (o output) isBot() bool {
	return false
}

type BotState struct {
	valueList    map[bot][]value
	movementList map[bot]*movement
}

func compileInstructions(instructions []string) (*BotState, error) {
	state := &BotState{}

	for _, v := range instructions {
		switch strings.Fields(v)[0] {
		case "value":
			botId, value := NewValue(v)
			if len(state.valueList[botId]) >= 2 {
				state.IterateCalcs()
			}
			state.valueList[botId] = append(state.valueList[botId], value)
		case "bot":
			botId, move := NewMovement(v)
			state.movementList[botId] = move
		}
	}
	return state, nil
}

func NewValue(val string) (bot, value) {
	fields := strings.Fields(val)
	botVal, err := strconv.Atoi(fields[5])
	if err != nil {
		panic(err)
	}
	valueVal, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	return bot(botVal), value(valueVal)
}

func NewMovement(val string) (bot, *movement) {
	fields := strings.Fields(val)
	botNum, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	var low, high Destination
	lowVal, err := strconv.Atoi(fields[6])
	if err != nil {
		panic(err)
	}
	switch fields[5] {
	case "bot":
		low = bot(lowVal)
	case "output":
		low = output(lowVal)
	}
	highVal, err := strconv.Atoi(fields[11])
	if err != nil {
		panic(err)
	}
	switch fields[10] {
	case "bot":
		high = bot(highVal)
	case "output":
		high = output(highVal)
	}
	return bot(botNum), &movement{
		lowTo:  low,
		highTo: high,
	}
}

func (b *BotState) IterateCalcs() {

}
