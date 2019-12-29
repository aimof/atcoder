package main

import (
	"testing"
)

func TestDo(t *testing.T) {
	testCases := []struct {
		in0  string
		in1  string
		want string
	}{}

	for i, tc := range testCases {
		got, err := do(tc.in0, tc.in1)
		if err != nil {
			t.Error()
		}
		if got != tc.want {
			t.Error(i, got, tc.want)
		}
	}
}
