package main

import (
	"bufio"
	"bytes"
	"strconv"

	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

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

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if ch == '\n' {
		return EOL
	}
	return ch
}

func NewScanner(line string) *Scanner {
	&Scanner{strings.NewReader(line)}
}

func decompress() {
	lines := helpers.MustScanFile("input.txt")
	for lines.Scan() {
		line := lines.Text()
	}
}
