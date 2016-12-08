package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
)

const doorID = "ugkcyxxp"

func HasFiveZeros(idx int, id string) bool {
	hash := md5.New()
	io.WriteString(hash, id)
	var counter int64
	var counterByte []byte
	binary.PutVarint(counterByte, counter)
	hash.Sum(counterByte)
	return false
}

func SumString(thing string) []byte {
	hash := md5.New()
	io.WriteString(hash, thing)
	return []byte{}
}

func GetPasswordForId(doorId string) string {
	return doorId
}

func main() {
	fmt.Printf("Door ID is %s, Password is %s", doorID, GetPasswordForId(doorID))
}
