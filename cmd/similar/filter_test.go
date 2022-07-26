package main

import (
	"fmt"
	"testing"
)

func TestNoneFilter(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"foobar", "foobar"},
		{"2022-07-22T10:57:00.000000+01:00 foo", "2022-07-22T10:57:00.000000+01:00 foo"},
		{"2022-07-22T10:57:00.000000Z foo", "2022-07-22T10:57:00.000000Z foo"},
		{"2022-07-22T10:57:00 baz", "2022-07-22T10:57:00 baz"},
		{"foo23", "foo23"},
		{"foo=1,bar=2", "foo=1,bar=2"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ret := noneFilter(tt.s)
			if ret != tt.want {
				t.Errorf("got %s, want %s", ret, tt.want)
			}
		})
	}
}

func TestExactFilter(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"foobar", "foobar"},
		{"2022-07-22T10:57:00.000000+01:00 foo", " foo"},
		{"2022-07-22T10:57:00.000000Z foo", " foo"},
		{"2022-07-22T10:57:00 baz", "2022-07-22T10:57:00 baz"},
		{"foo23", "foo23"},
		{"foo=1,bar=2", "foo=1,bar=2"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ret := exactFilter(tt.s)
			if ret != tt.want {
				t.Errorf("got %s, want %s", ret, tt.want)
			}
		})
	}
}

func TestNumbersFilter(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"foobar", "foobar"},
		{"2022-07-22T10:57:00.000000+01:00 foo", "--T::.+: foo"},
		{"2022-07-22T10:57:00.000000Z foo", "--T::.Z foo"},
		{"2022-07-22T10:57:00 baz", "--T:: baz"},
		{"foo23", "foo"},
		{"foo=1,bar=2", "foo=,bar="},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ret := numbersFilter(tt.s)
			if ret != tt.want {
				t.Errorf("got %s, want %s", ret, tt.want)
			}
		})
	}
}

func TestSignatureFilter(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"foobar", ""},
		{"2022-07-22T10:57:00.000000+01:00 foo", "--::.+: "},
		{"2022-07-22T10:57:00.000000Z foo", "--::. "},
		{"2022-07-22T10:57:00 baz", "--:: "},
		{"foo23", ""},
		{"foo=1,bar=2", "=,="},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ret := signatureFilter(tt.s)
			if ret != tt.want {
				t.Errorf("got %s, want %s", ret, tt.want)
			}
		})
	}
}
