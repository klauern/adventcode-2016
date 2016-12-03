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
		{
			"West Left Turn",
			fields{0, 0, 'N'},
			args{"L0"},
			'W',
		},
		{
			"West Right Turn",
			fields{0, 0, 'S'},
			args{"R0"},
			'W',
		},
		{
			"North Left Turn",
			fields{0, 0, 'E'},
			args{"L0"},
			'N',
		},
		{
			"North Right Turn",
			fields{0, 0, 'W'},
			args{"R0"},
			'N',
		},
		{
			"East Left Turn",
			fields{0, 0, 'S'},
			args{"L0"},
			'E',
		},
		{
			"East Right Turn",
			fields{0, 0, 'N'},
			args{"R0"},
			'E',
		},
		{
			"South Left Turn",
			fields{0, 0, 'W'},
			args{"L0"},
			'S',
		},
		{
			"South Right Turn",
			fields{0, 0, 'E'},
			args{"R0"},
			'S',
		},
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
		want   *Position
	}{
		{
			"L2 Facing N",
			fields{0, 0, 'N'},
			args{"L2"},
			&Position{X: -2, Y: 0, DirectionFacing: 'W'},
		},
		{
			"R2 Facing N",
			fields{0, 0, 'N'},
			args{"R2"},
			&Position{X: 2, Y: 0, DirectionFacing: 'E'},
		},
		{
			"R2 Facing E",
			fields{0, 0, 'E'},
			args{"R2"},
			&Position{X: 0, Y: -2, DirectionFacing: 'S'},
		},
		{
			"L2 Facing E",
			fields{0, 0, 'E'},
			args{"L2"},
			&Position{X: 0, Y: 2, DirectionFacing: 'N'},
		},
		{
			"L2 Facing W",
			fields{0, 0, 'W'},
			args{"L2"},
			&Position{X: 0, Y: -2, DirectionFacing: 'S'},
		},
		{
			"R2 Facing W",
			fields{0, 0, 'W'},
			args{"R2"},
			&Position{X: 0, Y: 2, DirectionFacing: 'N'},
		},
		{
			"R2 Facing S",
			fields{0, 0, 'S'},
			args{"R2"},
			&Position{X: -2, Y: 0, DirectionFacing: 'W'},
		},
		{
			"L2 Facing S",
			fields{0, 0, 'S'},
			args{"L2"},
			&Position{X: 2, Y: 0, DirectionFacing: 'E'},
		},
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

func TestPosition_String(t *testing.T) {
	tests := []struct {
		name string
		p    Position
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.p.String(); got != tt.want {
			t.Errorf("%q. Position.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCalculateMovement(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want *Position
	}{
		{
			"Two Turns",
			args{"R2, L3"},
			&Position{X: 2, Y: 3, DirectionFacing: 'N'},
		},
		{
			"Three Turns",
			args{"R2, R2, R2"},
			&Position{X: 0, Y: -2, DirectionFacing: 'W'},
		},
		{
			"Four Turns",
			args{"R5, L5, R5, R3"},
			&Position{X: 10, Y: 2, DirectionFacing: 'S'},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := CalculateMovement(tt.args.moves); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CalculateMovement() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
