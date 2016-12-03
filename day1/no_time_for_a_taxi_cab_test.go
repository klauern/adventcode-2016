package main

import (
	"reflect"
	"testing"
)

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

func TestPosition_CalculateDirection(t *testing.T) {
	type fields struct {
		X               int
		Y               int
		DirectionFacing rune
	}
	type args struct {
		next string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rune
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := Position{
			X:               tt.fields.X,
			Y:               tt.fields.Y,
			DirectionFacing: tt.fields.DirectionFacing,
		}
		if got := p.CalculateDirection(tt.args.next); got != tt.want {
			t.Errorf("%q. Position.CalculateDirection() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPosition_CalculatePosition(t *testing.T) {
	type fields struct {
		X               int
		Y               int
		DirectionFacing rune
	}
	type args struct {
		next string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := Position{
			X:               tt.fields.X,
			Y:               tt.fields.Y,
			DirectionFacing: tt.fields.DirectionFacing,
		}
		if got := p.CalculatePosition(tt.args.next); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Position.CalculatePosition() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
