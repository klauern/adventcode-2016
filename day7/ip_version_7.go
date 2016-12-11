package main

type IPv7 struct {
	nonHyperNetSeqs []string // even sequences (0, 2, 4, 6, ...)
	hyperNetSeqs    []string // odd sequences (1, 3, 5, 7, ...)
	supportsTLS     bool
}

// hasABBA calculates whether the given string has an ABBA.
//
// An ABBA is any four-character sequence which consists of a pair of two different
// characters followed by the reverse of that pair, such as xyyx or abba.
func hasABBA(str string) bool {
	if len(str) < 4 {
		return false
	}
	for i := 1; i < len(str)-1; i++ {
		if str[i] == str[i-1] {
			continue
		}
		if str[i] == str[i+1] && str[i-1] == str[i+2] {
			return true
		}
	}
	return false
}

func NewIPv7(str string) IPv7 {

}
