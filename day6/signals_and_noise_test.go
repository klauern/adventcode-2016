package main

import (
	"reflect"
	"testing"
)

func Test_code_Decode(t *testing.T) {
	tests := []struct {
		name string
		c    code
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.DecodeMostCommon(); got != tt.want {
			t.Errorf("%q. code.Decode() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_code_SliceToColumns(t *testing.T) {
	tests := []struct {
		name string
		c    code
		want code
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.SliceToColumns(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. code.SliceToColumns() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_countRuneAry(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := countRuneAry(tt.args.chars); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. countRuneAry() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_commonRune(t *testing.T) {
	type args struct {
		runeCount map[rune]int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mostCommonRune(tt.args.runeCount); got != tt.want {
			t.Errorf("%q. commonRune() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findCommonRune(t *testing.T) {
	type args struct {
		runes []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findMostCommonRune(tt.args.runes); got != tt.want {
			t.Errorf("%q. findCommonRune() = %v, want %v", tt.name, got, tt.want)
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

func TestNewCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want code
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewCode(tt.args.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewCode() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDecodeMostCommonString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"thing",
			args{
				`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`,
			},
			"easter",
		},
	}
	for _, tt := range tests {
		if got := DecodeMostCommonString(tt.args.s); got != tt.want {
			t.Errorf("%q. DecodeString() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDecodeLeastCommonString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"thing",
			args{
				`eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`,
			},
			"advent",
		},
	}
	for _, tt := range tests {
		if got := DecodeLeastCommonString(tt.args.s); got != tt.want {
			t.Errorf("%q. DecodeLeastCommonString() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_code_DecodeLeastCommon(t *testing.T) {
	tests := []struct {
		name string
		c    code
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.DecodeLeastCommon(); got != tt.want {
			t.Errorf("%q. code.DecodeLeastCommon() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_mostCommonRune(t *testing.T) {
	type args struct {
		runeCount map[rune]int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mostCommonRune(tt.args.runeCount); got != tt.want {
			t.Errorf("%q. mostCommonRune() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findMostCommonRune(t *testing.T) {
	type args struct {
		runes []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findMostCommonRune(tt.args.runes); got != tt.want {
			t.Errorf("%q. findMostCommonRune() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findLeastCommonRune(t *testing.T) {
	type args struct {
		runes []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findLeastCommonRune(tt.args.runes); got != tt.want {
			t.Errorf("%q. findLeastCommonRune() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_leastCommonRune(t *testing.T) {
	type args struct {
		runeCount map[rune]int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := leastCommonRune(tt.args.runeCount); got != tt.want {
			t.Errorf("%q. leastCommonRune() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
