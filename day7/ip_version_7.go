package main

import "strings"

// IPv7 represents Day 7's Internet Protocol version 7 specification.
// It consists of pairs of hypernet sequences (surrounded by []'s) with
// non-hyperNet sequences.
type IPv7 struct {
	nonHyperNetSeqs []string // even sequences (0, 2, 4, 6, ...)
	hyperNetSeqs    []string // odd sequences (1, 3, 5, 7, ...)
	isTLSSupported  bool
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

func NewIPv7(str string) *IPv7 {
	ip := &IPv7{}
	sections := strings.Split(str, "]")
	for _, section := range sections {
		nets := strings.Split(section, "[")
		ip.nonHyperNetSeqs = append(ip.nonHyperNetSeqs, nets[0])
		if len(nets) > 1 {
			ip.hyperNetSeqs = append(ip.hyperNetSeqs, nets[1])
		}
	}
	ip.isTLSSupported = checkTLSSupported(ip.nonHyperNetSeqs, ip.hyperNetSeqs)
	return ip
}

func checkTLSSupported(non, hyper []string) bool {
	for _, val := range hyper {
		if hasABBA(val) {
			return false
		}
	}
	for _, val := range non {
		if hasABBA(val) {
			return true
		}
	}
	return false
}
