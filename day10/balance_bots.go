package main

import (
	"strconv"
	"strings"
)

const (
	botType    = "bot"
	valType    = "value"
	outputType = "output"
)

type (
	value  int
	botID  int
	output int
	bot    struct {
		id   botID
		vals []value
	}
	movement struct {
		lowTo  Destination
		highTo Destination
	}
)

// BotState represents the entire state of the assembly line in a particular
// point in time.
type BotState struct {
	bots         map[botID]bot
	movementList map[botID]*movement
	outputList   map[output][]value
}

func compileInstructions(instructions []string) (*BotState, error) {
	state := &BotState{}

	for _, v := range instructions {
		switch strings.Fields(v)[0] {
		case valType:
			id, chip := NewValue(v)
			if state.bots[id].canReceive() {
				state.bots[botID].vals = append(state.bots[botID].vals, chip)
			} else {
				state.IterateCalcs()
			}
		case botType:
			botID, move := NewMovement(v)
			state.movementList[botID] = move
		}
	}
	return state, nil
}

// NewValue creates a new chip assignment to a bot.
func NewValue(val string) (botID, value) {
	fields := strings.Fields(val)
	botVal, err := strconv.Atoi(fields[5])
	if err != nil {
		panic(err)
	}
	valueVal, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	return botID(botVal), value(valueVal)
}

// NewMovement will parse an instruction line and create a new
// bot and movement rule.
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
	case botType:
		low = bot(lowVal)
	case outputType:
		low = output(lowVal)
	}
	highVal, err := strconv.Atoi(fields[11])
	if err != nil {
		panic(err)
	}
	switch fields[10] {
	case botType:
		high = bot(highVal)
	case outputType:
		high = output(highVal)
	}
	return bot(botNum), &movement{
		lowTo:  low,
		highTo: high,
	}
}

// IterateCalcs will iterate over the BotState and determine the calculations
// for all bots with two chips.
func (b *BotState) IterateCalcs() {
	for k, v := range b.valueList {
		if len(v) == 2 {
			movement := b.movementList[k]
			b = b.move(k)
		}
	}

}

func (b *BotState) move(id bot) {
	values := b.valueList[id]
	var high, low value
	if values[0] > values[1] {
		high, low = values[0], values[1]
	} else {
		high, low = values[1], values[0]
	}
	movement := b.movementList[id]

	// this is likely recursive, as you will be moving the high/low vals to respective
	// bots.  You start with the low, then move the high, recursing on this same call
	// to make calls if the destination that you're going to need to do i
}
