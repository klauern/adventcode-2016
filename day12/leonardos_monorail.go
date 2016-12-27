package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

type Register int

type Program struct {
	counter      int
	instructions []string
	aReg         *Register
	bReg         *Register
	cReg         *Register
	dReg         *Register
}

var logger *log.Logger

func isRegister(val string) bool {
	return val == "a" || val == "b" || val == "c" || val == "d"
}

func (p *Program) execute(instr []string) bool {
	switch instr[0] {
	case "cpy":
		p.copy(instr[1], instr[2])
	case "inc":
		p.increment(instr[1])
	case "dec":
		p.decrement(instr[1])
	case "jnz":
		return p.jumpNZ(instr[1], instr[2])
	}
	return false
}

func (p *Program) Run() {
	for p.counter < len(p.instructions) {
		fields := strings.Fields(p.instructions[p.counter])
		cont := p.execute(fields)
		logger.Printf("counter: %v", p.counter)
		logger.Printf("%v", p)
		if cont {
			continue
		}
		p.counter++
	}
}

func (p *Program) getRegister(reg string) *Register {
	switch reg {
	case "a":
		return p.aReg
	case "b":
		return p.bReg
	case "c":
		return p.cReg
	case "d":
		return p.dReg
	}
	logger.Printf("Nil value retrieved for Register %s\n", reg)
	return nil
}

func (p *Program) copy(from, to string) *Register {
	toReg := p.getRegister(to)
	if isRegister(from) {
		fVal := p.getRegister(from)
		if fVal != nil {
			*toReg = *fVal
		}
	} else {
		fVal, err := strconv.Atoi(from)
		if err != nil {
			panic(err)
		}
		r := Register(fVal)
		toReg = &r
	}
	return toReg
}

func (p *Program) increment(register string) *Register {
	reg := p.getRegister(register)
	*reg++
	return reg
}

func (p *Program) decrement(register string) *Register {
	reg := p.getRegister(register)
	*reg--
	return reg
}

func (p *Program) jumpNZ(reg, delta string) bool {
	if isRegister(reg) {
		r := p.getRegister(reg)
		if r != nil {
			logger.Printf("register %d\n", *r)
			if int(*r) != 0 {
				p.jump(delta)
				return true
			}
		}
		return false
	}
	val, err := strconv.Atoi(reg)
	if err != nil {
		panic(err)
	}
	if val != 0 {
		p.jump(delta)
		return true
	}
	return false
}

func (p *Program) jump(delta string) {
	jump, err := strconv.Atoi(delta)
	if err != nil {
		panic(err)
	}
	p.counter += jump
}

func (p *Program) String() string {
	return fmt.Sprintf("Program Register values: a: %v b: %v c: %v d: %v", *p.aReg, *p.bReg, *p.cReg, *p.dReg)
}

func main() {
	logger = log.New(os.Stdout, "", 0)
	file := helpers.MustScanFile("input.txt")
	instructionSet := make([]string, 0)
	for file.Scan() {
		line := file.Text()
		fmt.Printf("Line %v\n", line)
		instructionSet = append(instructionSet, line)
	}
	a := Register(0)
	b := Register(0)
	c := Register(0)
	d := Register(0)
	program := &Program{instructions: instructionSet, aReg: &a, bReg: &b, cReg: &c, dReg: &d}
	program.Run()
	fmt.Println(program)
}
