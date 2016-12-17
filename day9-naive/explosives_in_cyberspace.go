package main

import (
	"regexp"
	"strconv"
	"strings"
)

type multiplier struct {
	charCount,
	times int
}

var mults, nonMults *regexp.Regexp

func init() {
	mults = regexp.MustCompile(`\([0-9]+x[0-9]+\)`)
	nonMults = regexp.MustCompile(`\)?[:alpha:]?+\(`)
}

func splitMultipliers(s string) []*multiplier {
	mstrs := mults.FindAllString(s, -1)
	multipliers := make([]*multiplier, 0)
	for _, v := range mstrs {
		multipliers = append(multipliers, NewMultiplier(v))
	}
	return multipliers
}

func splitNonMultipliers(s string) []string {
	return nonMults.FindAllString(s, -1)
}

func NewMultiplier(s string) *multiplier {
	m := strings.TrimPrefix(s, "(")
	m = strings.TrimSuffix(m, ")")
	parts := strings.Split(m, "x")
	chars, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	counts, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return &multiplier{
		charCount: chars,
		times:     counts,
	}
}
