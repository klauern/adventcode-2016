package helpers

import (
	"bufio"
	"io/ioutil"
	"os"
)

func MustLoadFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func MustScanFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}
