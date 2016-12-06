package main

import (
	"reflect"
	"testing"
)

func TestKeyPos_CalcMoves(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		move rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   KeyPos
	}{
		{
			"Move Left",
			fields{1, 1},
			args{'L'},
			KeyPos{0, 1},
		},
		{
			"Move Up",
			fields{1, 1},
			args{'U'},
			KeyPos{1, 0},
		},
		{
			"Move Down",
			fields{1, 1},
			args{'D'},
			KeyPos{1, 2},
		},
		{
			"Move Right",
			fields{1, 1},
			args{'R'},
			KeyPos{2, 1},
		},
		{
			"Move Left (No space)",
			fields{0, 1},
			args{'L'},
			KeyPos{0, 1},
		},
		{
			"Move Right (No space)",
			fields{2, 1},
			args{'R'},
			KeyPos{2, 1},
		},
		{
			"Move Up (no space)",
			fields{1, 0},
			args{'U'},
			KeyPos{1, 0},
		},
		{
			"Move Down (no space)",
			fields{1, 2},
			args{'D'},
			KeyPos{1, 2},
		},
	}
	for _, tt := range tests {
		k := &KeyPos{
			X: tt.fields.X,
			Y: tt.fields.Y,
		}
		if got := k.CalcMoves(tt.args.move); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. KeyPos.CalcMoves() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestKeyPos_CalcNextDigit(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		moves string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   KeyPos
	}{
		{
			"ULL (1)",
			fields{1, 1},
			args{"ULL"},
			KeyPos{0, 0},
		},
		{
			"RRDDD (9)",
			fields{0, 0},
			args{"RRDDD"},
			KeyPos{2, 2},
		},
		{
			"LURDL (8)",
			fields{2, 2},
			args{"LURDL"},
			KeyPos{1, 2},
		},
		{
			"UUUUD (5)",
			fields{1, 1},
			args{"UUUUD"},
			KeyPos{1, 1},
		},
	}
	for _, tt := range tests {
		k := &KeyPos{
			X: tt.fields.X,
			Y: tt.fields.Y,
		}
		if got := k.CalcNextDigit(tt.args.moves); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. KeyPos.CalcNextDigit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}

func TestKeyPos_GetDigit(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"1",
			fields{0, 0},
			1,
		},
		{
			"2",
			fields{1, 0},
			2,
		},
		{
			"3",
			fields{2, 0},
			3,
		},
		{
			"4",
			fields{0, 1},
			4,
		},
		{
			"5",
			fields{1, 1},
			5,
		},
		{
			"6",
			fields{2, 1},
			6,
		},
		{
			"7",
			fields{0, 2},
			7,
		},
		{
			"8",
			fields{1, 2},
			8,
		},
		{
			"9",
			fields{2, 2},
			9,
		},
	}
	for _, tt := range tests {
		k := &KeyPos{
			X: tt.fields.X,
			Y: tt.fields.Y,
		}
		if got := k.GetDigit(); got != tt.want {
			t.Errorf("%q. KeyPos.GetDigit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
