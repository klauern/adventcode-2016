package main

import (
	"reflect"
	"testing"
)

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

func TestNewIPv7(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *IPv7
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewIPv7(tt.args.str); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewIPv7() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIPv7_supportsTLS(t *testing.T) {
	type fields struct {
		nonHyperNetSeqs []string
		hyperNetSeqs    []string
		isTLSSupported  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"abba[mnop]qrst",
			fields{
				[]string{"abba", "qrst"},
				[]string{"mnop"},
				true,
			},
			true,
		},
	}
	for _, tt := range tests {
		ip := &IPv7{
			nonHyperNetSeqs: tt.fields.nonHyperNetSeqs,
			hyperNetSeqs:    tt.fields.hyperNetSeqs,
			isTLSSupported:  tt.fields.isTLSSupported,
		}
		if got := ip.supportsTLS(); got != tt.want {
			t.Errorf("%q. IPv7.supportsTLS() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
