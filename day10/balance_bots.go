package main

import (
	"strconv"
	"strings"
)

type valueMove struct {
	value,
	bot int
}

type movement struct {
	high,
	low int
}

type value int
type bot int

type InstructionSet struct {
	valueList    map[value]bot
	movementList []movement
}

func compileInstructions(instructions []string) (*InstructionSet, error) {

	for _, v := range instructions {
		if strings.HasPrefix(v, "value") {

		}
	}
	return nil, nil
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
