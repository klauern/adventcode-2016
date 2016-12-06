package main

import (
	"reflect"
	"testing"
)

func TestLargestSide(t *testing.T) {
	type args struct {
		sides []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"5 10 25",
			args{[]float64{5, 10, 25}},
			25,
		},
		{
			"50 10 25",
			args{[]float64{50, 10, 25}},
			50,
		},
		{
			"5 100 25",
			args{[]float64{5, 100, 25}},
			100,
		},
	}
	for _, tt := range tests {
		if got := LargestSide(tt.args.sides); got != tt.want {
			t.Errorf("%q. LargestSide() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_canTriangle(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"25, 5, 10 (no)",
			args{25, 5, 10},
			false,
		},
	}
	for _, tt := range tests {
		if got := canTriangle(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
			t.Errorf("%q. canTriangle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIsTriangle(t *testing.T) {
	type args struct {
		sides []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"5 10 25 (no)",
			args{[]float64{5, 10, 25}},
			false,
		},
	}
	for _, tt := range tests {
		if got := IsTriangle(tt.args.sides); got != tt.want {
			t.Errorf("%q. IsTriangle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewSide(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"4   5  10",
			args{"4   5  10"},
			[]float64{4.0, 5.0, 10.0},
		},
	}
	for _, tt := range tests {
		if got := NewSide(tt.args.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewSide() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
