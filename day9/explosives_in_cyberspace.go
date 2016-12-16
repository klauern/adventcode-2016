package main

import (
	"bytes"
	"strconv"

	"github.com/klauern/adventcode-2016/helpers"
)

type Repeating struct {
	charCount,
	times int
}

func (s *Scanner) readMultiplier() *Repeating {
	var buf bytes.Buffer
	s.read()
	var charCount, times int
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == 'X' {
			cnt, err := strconv.Atoi(buf.String())
			if err != nil {
				charCount = -1
			} else {
				charCount = cnt
			}
			buf.Reset()

		} else if ch == ')' {
			cnt, err := strconv.Atoi(buf.String())
			if err != nil {
				times = -1
			} else {
				times = cnt
			}
			buf.Reset()
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	return &Repeating{charCount, times}
}

func decompress() {
	lines := helpers.MustScanFile("input.txt")
	for lines.Scan() {
		line := lines.Text()
	}
}
