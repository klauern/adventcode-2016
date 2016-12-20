package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

const (
	botType    = "bot"
	valType    = "value"
	outputType = "output"
)

type Destination interface {
	isOutput() bool
	isBot() bool
	canReceive() bool
	receiveValue(value)
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

func (o output) canReceive() bool {
	return true
}

func (b bot) canReceive() bool {
	return len(b.vals) < 2
}

func (o output) receiveValue(val value) {
	o.vals = append(o.vals, val)
}

func (b bot) receiveValue(val value) {
	b.vals = append(b.vals, val)
}

type (
	value    int
	botID    int
	outputID int
	output   struct {
		id   outputID
		vals []value
	}
	bot struct {
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
	bots         map[botID]*bot
	movementList map[botID]*movement
	outputList   map[int][]value
}

func compileInstructions(file *bufio.Scanner) (*BotState, error) {
	state := &BotState{}
	for file.Scan() {
		line := file.Text()
		switch strings.Fields(line)[0] {
		case valType:
			id, chip := NewValue(line)
			if state.bots[id].canReceive() {
				state.bots[id].vals = append(state.bots[id].vals, chip)
			} else {
				state.IterateCalcs()
			}
		case botType:
			b, move := NewMovement(line)
			state.movementList[b] = move
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
func NewMovement(val string) (botID, *movement) {
	fields := strings.Fields(val)
	botNum, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	var lowDest, highDest Destination
	lowVal, err := strconv.Atoi(fields[6])
	if err != nil {
		panic(err)
	}
	switch fields[5] {
	case botType:
		lowDest = bot{id: botID(-1), vals: []value{value(lowVal)}}
	case outputType:
		lowDest = output{vals: []value{value(lowVal)}}
	}
	highVal, err := strconv.Atoi(fields[11])
	if err != nil {
		panic(err)
	}
	switch fields[10] {
	case botType:
		highDest = bot{id: botID(highVal)}
	case outputType:
		highDest = output{id: outputID(highVal)}
	}
	return botID(botNum), &movement{
		lowTo:  lowDest,
		highTo: highDest,
	}
}

// IterateCalcs will iterate over the BotState and determine the calculations
// for all bots with two chips.
func (b *BotState) IterateCalcs() {
	for k, v := range b.bots {
		if len(v.vals) == 2 {
			movement := b.movementList[k]
			b.move(k)
		}
	}
}

func (b *BotState) move(id botID) {
	values := b.bots[id].vals
	var high, low value
	if values[0] > values[1] {
		high, low = values[0], values[1]
	} else {
		high, low = values[1], values[0]
	}
	movement := b.movementList[id]
	if movement.highTo.canReceive() {

	}

	// this is likely recursive, as you will be moving the high/low vals to respective
	// bots.  You start with the low, then move the high, recursing on this same call
	// to make calls if the destination that you're going to need to do i
}

func main() {
	file := helpers.MustScanFile("input.txt")
	state, err := compileInstructions(file)
	if err != nil {
		panic(err)
	}

	for _, v := range state.bots {
		if len(v.vals) > 1 {
			fmt.Printf("Bot %v", v)
		}
	}
}
