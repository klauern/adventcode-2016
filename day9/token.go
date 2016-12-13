package main

type Token int

const (
	EOF = iota
	EOL
	WS

	START_MULTIPLIER
	END_MULTIPLIER
	X_MULTIPLIER

	CHARACTER
)
