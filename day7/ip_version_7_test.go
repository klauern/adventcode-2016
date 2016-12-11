package main

import (
	"reflect"
	"testing"
)

func Test_hasABBA(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abba : true",
			args{"abba"},
			true,
		},
		{
			"abcd: false",
			args{"abcd"},
			false,
		},
		{
			"aaaa: false",
			args{"aaaa"},
			false,
		},
		{
			"ioxxoj: true",
			args{"ioxxoj"},
			true,
		},
	}
	for _, tt := range tests {
		if got := hasABBA(tt.args.str); got != tt.want {
			t.Errorf("%q. hasABBA() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewIPv7(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *IPv7
	}{
		{
			"abba[mnop]qrst",
			args{"abba[mnop]qrst"},
			&IPv7{
				hyperNetSeqs:    subsequence{"mnop"},
				nonHyperNetSeqs: subsequence{"abba", "qrst"},
				isTLSSupported:  true,
			},
		},
		{
			"abcd[bddb]xyyx",
			args{"abcd[bddb]xyyx"},
			&IPv7{
				hyperNetSeqs:    subsequence{"bddb"},
				nonHyperNetSeqs: subsequence{"abcd", "xyyx"},
				isTLSSupported:  false,
			},
		},
		{
			"aaaa[qwer]tyui",
			args{"aaaa[qwer]tyui"},
			&IPv7{
				hyperNetSeqs:    subsequence{"qwer"},
				nonHyperNetSeqs: subsequence{"aaaa", "tyui"},
				isTLSSupported:  false,
			},
		},
		{
			"ioxxoj[asdfgh]zxcvbn",
			args{"ioxxoj[asdfgh]zxcvbn"},
			&IPv7{
				hyperNetSeqs:    subsequence{"asdfgh"},
				nonHyperNetSeqs: subsequence{"ioxxoj", "zxcvbn"},
				isTLSSupported:  true,
			},
		},
	}
	for _, tt := range tests {
		if got := NewIPv7(tt.args.str); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewIPv7() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_subsequence_hasABBA(t *testing.T) {
	tests := []struct {
		name string
		s    subsequence
		want bool
	}{
		{
			"abba, qrst",
			subsequence{"abba", "qrst"},
			true,
		},
		{
			"mnop",
			subsequence{"mnop"},
			false,
		},
		{
			"aaaa, tyui",
			subsequence{"aaaa", "tyui"},
			false,
		},
	}
	for _, tt := range tests {
		if got := tt.s.hasABBA(); got != tt.want {
			t.Errorf("%q. subsequence.hasABBA() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
