package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetElement(t *testing.T) {
	var tests = []struct {
		arr     []int
		value   int
		want    int
		errWant string
	}{
		{
			[]int{},
			3,
			0,
			"empty slice",
		},
		{
			[]int{1, 2, 3, 4, 5},
			-1,
			0,
			"invalid idx",
		},
		{
			[]int{1, 2, 3, 4, 5},
			6,
			0,
			"invalid idx",
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			3,
			"",
		},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("getElement_test#%d", idx+1)
		t.Run(testname, func(t *testing.T) {
			res, err := getElem(tt.arr, tt.value)

			if res != tt.want {
				t.Error("invalid test")
			}

			if err != nil && len(tt.arr) == 0 && (!strings.Contains(err.Error(), tt.errWant) || tt.want != res) {
				t.Error("invalid test")
			}

			if err != nil && (tt.value < 0 || tt.value > len(tt.arr)) && (!strings.Contains(err.Error(), tt.errWant) || tt.want != res) {
				t.Error("invalid test")
			}
		})
	}

}
