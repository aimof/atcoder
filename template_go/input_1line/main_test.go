package main

import (
	"testing"
)

func TestDo(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{}

	for i, tc := range testCases {
		got, err := do(tc.in)
		if err != nil {
			t.Error()
		}
		if got != tc.want {
			t.Errorf("case: %d\ngot:  %v\nwant: %v", i, got, tc.want)
		}
	}
}
