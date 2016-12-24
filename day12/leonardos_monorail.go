package main

import (
	"fmt"
	"strconv"
	"strings"

	"log"
	"os"

	"github.com/klauern/adventcode-2016/helpers"
)

type Instruction int

var logger *log.Logger

const (
	cpy Instruction = iota
	inc
	dec
	jnz
	unknown
)

type Register int

type Program struct {
	counter      int
	instructions []string
	aReg         Register
	bReg         Register
	cReg         Register
	dReg         Register
}

func isRegister(val string) bool {
	return val == "a" || val == "b" || val == "c" || val == "d"
}

func (p *Program) execute(instr []string) bool {
	logger.Printf("switch on %v\n", instr[0])
	switch instr[0] {
	case "cpy":
		p.copy(instr[1], instr[2])
	case "inc":
		p.increment(instr[1])
	case "dec":
		p.decrement(instr[1])
	case "jnz":
		jumped := p.jumpNZ(instr[1], instr[2])
		return jumped
	}
	return false
}

func (p *Program) Run() {
	for p.counter < len(p.instructions) {
		fields := strings.Fields(p.instructions[p.counter])
		cont := p.execute(fields)
		if cont {
			continue
		}
		p.counter++
		logger.Printf("counter: %v", p.counter)
	}
}

func (p *Program) getRegister(reg string) *Register {
	switch reg {
	case "a":
		return &p.aReg
	case "b":
		return &p.bReg
	case "c":
		return &p.cReg
	case "d":
		return &p.dReg
	}
	return nil
}

func (p *Program) copy(from, to string) {
	toReg := p.getRegister(to)
	if isRegister(from) {
		fVal := p.getRegister(from)
		*toReg = *fVal
		logger.Printf("value of %s is %v\n", from, *toReg)
	} else {
		fVal, err := strconv.Atoi(from)
		if err != nil {
			panic(err)
		}
		*toReg = Register(fVal)
	}
}

func (p *Program) increment(register string) {
	reg := p.getRegister(register)
	*reg++
}

func (p *Program) decrement(register string) {
	reg := p.getRegister(register)
	*reg--
}

func (p *Program) jumpNZ(reg, delta string) bool {
	if isRegister(reg) {
		r := p.getRegister(reg)
		logger.Printf("register %v\n", *r)
		if int(*r) != 0 {
			jump, err := strconv.Atoi(delta)
			if err != nil {
				panic(err)
			}
			p.counter += jump
			return true
		}
		return false
	} else {
		val, err := strconv.Atoi(reg)
		if err != nil {
			panic(err)
		}
		if val != 0 {
			jump, err := strconv.Atoi(delta)
			if err != nil {
				panic(err)
			}
			p.counter += jump
			return true
		}
		return false
	}
}

func (p *Program) String() string {
	return fmt.Sprintf("Program Register values:\na: %v\nb: %v\nc: %v\nd: %v", p.aReg, p.bReg, p.cReg, p.dReg)
}

func main() {
	logger = log.New(os.Stdout, "day12 ", 0)
	file := helpers.MustScanFile("input.txt")
	instructionSet := make([]string, 0)
	for file.Scan() {
		line := file.Text()
		fmt.Printf("Line %v\n", line)
		instructionSet = append(instructionSet, line)
	}
	program := &Program{instructions: instructionSet}
	program.Run()
	fmt.Println(program)
}
