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

func TestPosition_CalculatePosition(t *testing.T) {
	type fields struct {
		North           int
		South           int
		East            int
		West            int
		DirectionFacing string
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
			North:           tt.fields.North,
			South:           tt.fields.South,
			East:            tt.fields.East,
			West:            tt.fields.West,
			DirectionFacing: tt.fields.DirectionFacing,
		}
		if got := p.CalculatePosition(tt.args.next); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Position.CalculatePosition() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
