package main

import (
	"strconv"
	"strings"
)

type Room struct {
	maxRuneField map[string]rune
	sector       int
	checksum     string
}

func (r *Room) CalcChecksum() {

}

func NewRoom(line string) *Room {
	fields := strings.Split(line, "-")
	room := &Room{}
	for _, v := range fields[:len(fields)-1] {
        maxRuneField[v] =
	}
	return room
}

func calcMax

func splitSectorChecksum(s string) (int, string) {
	secksum := strings.Split(s, "[")
	cksum := strings.Trim(secksum[1], "]")
	room, err := strconv.Atoi(secksum[0])
	if err != nil {
		panic(err)
	}
	return room, cksum
}

func (r *Room) RealRoom() bool {
	return false
}

func main() {

}
