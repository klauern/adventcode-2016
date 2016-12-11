package main

import "testing"

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
