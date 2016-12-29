package main

import (
	"reflect"
	"testing"
)

func TestChecksum(t *testing.T) {
	type args struct {
		room string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"aaaaa-bbb-z-y-x-123[abxyz]",
			args{"aaaaa-bbb-z-y-x-123[abxyz]"},
			"abxyz",
		},
		{
			"a-b-c-d-e-f-g-h-987[abcde]",
			args{"a-b-c-d-e-f-g-h-987[abcde]"},
			"abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum(tt.args.room); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountString(t *testing.T) {
	type args struct {
		room string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			"aaaabbbzyx",
			args{"aaaabbbzyx"},
			map[rune]int{'a': 4, 'b': 3, 'z': 1, 'y': 1, 'x': 1},
		},
		{
			"abcdefgh",
			args{"abcdefgh"},
			map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountString(tt.args.room); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	type args struct {
		maps []map[rune]int
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			"abcd, abd",
			args{
				[]map[rune]int{
					{'a': 1, 'b': 1, 'c': 1, 'd': 1},
					{'a': 1, 'b': 1, 'd': 1},
				},
			},
			map[rune]int{'a': 2, 'b': 2, 'c': 1, 'd': 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMaps(tt.args.maps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
