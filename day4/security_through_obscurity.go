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

func MergeMaps(maps []map[rune]int) map[rune]int {
	results := make(map[rune]int, 0)
	for _, m := range maps {
		for k, v := range m {
			results[k] += v
		}
	}
	return results
}
