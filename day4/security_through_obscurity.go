package main

import "strings"
import "sort"

type Room struct {
	sectors  []string
	checksum string
}

type RuneMap struct {
	runeCount map[rune]int
	runeAry   []rune
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

type ByRuneCount RuneMap

func (r ByRuneCount) Len() int           { return len(r.runeAry) }
func (r ByRuneCount) Swap(i, j int)      { r.runeAry[i], r.runeAry[j] = r.runeAry[j], r.runeAry[i] }
func (r ByRuneCount) Less(i, j int) bool { return r.runeCount[r.runeAry[i]] < r.runeCount[r.runeAry[j]] }

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

func CalcCheckSum(runeMap map[rune]int) []rune {
	var keys []rune
	for k := range runeMap {
		keys = append(keys, k)
	}
	sort.Sort(ByRune(keys))

	return keys
}
