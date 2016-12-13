package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestScanner_readMultiplier(t *testing.T) {
	type fields struct {
		r *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   *Repeating
	}{
		{
			"(3x3)",
			fields{bufio.NewReader(strings.NewReader("(3X3)"))},
			&Repeating{3, 3},
		},
		{
			"(7X10)",
			fields{bufio.NewReader(strings.NewReader("(7X10)"))},
			&Repeating{7, 10},
		},
	}
	for _, tt := range tests {
		s := &Scanner{
			r: tt.fields.r,
		}
		if got := s.readMultiplier(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Scanner.readMultiplier() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestScanner_read(t *testing.T) {
	type fields struct {
		r *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Scanner{
			r: tt.fields.r,
		}
		if got := s.read(); got != tt.want {
			t.Errorf("%q. Scanner.read() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
