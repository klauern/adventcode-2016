package main

import (
	"fmt"
	"strings"

	"github.com/klauern/adventcode-2016/helpers"
)

type subsequence []string

// IPv7 represents Day 7's Internet Protocol version 7 specification.
// It consists of pairs of hypernet sequences (surrounded by []'s) with
// non-hyperNet sequences.
type IPv7 struct {
	nonHyperNetSeqs subsequence // even sequences (0, 2, 4, 6, ...)
	hyperNetSeqs    subsequence // odd sequences (1, 3, 5, 7, ...)
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

func (s subsequence) hasABBA() bool {
	for _, val := range s {
		if hasABBA(val) {
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
	if ip.hyperNetSeqs.hasABBA() {
		ip.isTLSSupported = false
	} else {
		ip.isTLSSupported = ip.nonHyperNetSeqs.hasABBA()
	}
	return ip
}

func main() {
	amount := 0
	ipAddrs := helpers.MustLoadFile("input.txt")
	for _, v := range strings.Split(ipAddrs, "\n") {
		ipv7 := NewIPv7(v)
		if ipv7.isTLSSupported {
			amount++
		}
	}
	fmt.Printf("Number of supported TLS addresses: %v", amount)
}
