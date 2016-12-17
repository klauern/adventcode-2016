package main

type Token int

const (
	// Special tokens
	EOF Token = iota
	EOL
	WS
	ILLEGAL

	// multiplier chars
	START_MULTIPLIER
	END_MULTIPLIER
	X_MULTIPLIER // only useful when between DIGIT and START_MULTIPLIER or END_MULTIPLIER
	DIGIT

	// everything else
	CHARACTER
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
