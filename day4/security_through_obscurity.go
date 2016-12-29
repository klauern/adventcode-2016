package main

import "strings"

type Room struct {
	sectors  []string
	checksum string
}

func Checksum(room string) string {
	strs := strings.Split(room, "[")
	return strings.TrimRight(strs[1], "]")
}

func CountString(room string) map[rune]int {
	charMap := make(map[rune]int, 0)
	for _, char := range room {
		charMap[char]++
	}
	return charMap
}
