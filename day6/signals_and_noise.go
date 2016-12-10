package main

import (
	"fmt"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

type code [][]rune

func (c code) Decode() string {
	decoded := make([]rune, 0)
	colCode := c.SliceToColumns()
	for i := range colCode {
		cnts := countRuneAry(colCode[i])
		common := commonRune(cnts)
		decoded = append(decoded, common)
	}
	return string(decoded)
}

func (c code) SliceToColumns() code {
	columns := make(code, 0)
	for i := 0; i < len(c[0]); i++ {
		column := make([]rune, 0)
		for _, row := range c {
			column = append(column, row[i])
		}
		columns = append(columns, column)
	}
	return columns
}

func countRuneAry(chars []rune) map[rune]int {
	counts := make(map[rune]int)
	for _, v := range chars {
		counts[v]++
	}
	return counts
}

func commonRune(runeCount map[rune]int) rune {
	highestVal := -1
	var highestRune rune
	for k, v := range runeCount {
		if runeCount[k] > highestVal {
			highestVal = v
			highestRune = k
		}
	}
	return highestRune
}

func findCommonRune(runes []rune) rune {
	counts := countRuneAry(runes)
	return commonRune(counts)
}

func NewCode(s string) code {
	c := make(code, 0)
	strAry := strings.Split(s, "\n")
	for i := range strAry {
		c = append(c, []rune(strAry[i]))
	}
	return c
}

func main() {
	str := helpers.MustLoadFile("input.txt")
	c := NewCode(str)
	fmt.Println(c.Decode())
}
