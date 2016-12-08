package main

import (
	"reflect"
	"testing"
)

func TestSumString(t *testing.T) {
	type args struct {
		thing string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := SumString(tt.args.thing); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SumString() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetPasswordForId(t *testing.T) {
	type args struct {
		doorId string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"abc",
			args{"abc"},
			"18f47a30",
		},
	}
	for _, tt := range tests {
		if got := GetPasswordForId(tt.args.doorId); got != tt.want {
			t.Errorf("%q. GetPasswordForId(\"%v\") = %v, want %v", tt.args.doorId, tt.name, got, tt.want)
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
