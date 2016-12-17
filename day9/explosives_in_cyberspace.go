package main

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

type Repeating struct {
	charCount,
	times int
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
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
