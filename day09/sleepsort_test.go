package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSleepSort(t *testing.T) {
	var tests = []struct {
		value []int
		want  []int
	}{
		{
			value: []int{5, 4, 3, 2, 1, 0},
			want:  []int{0, 1, 2, 3, 4, 5},
		},
		{
			value: []int{},
			want:  []int{},
		},
		{
			value: []int{4, 2, 1, 9, 6, 7, 3, 5, 8, 0},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("TestSleepSort#%d", idx+1)

		got := make([]int, 0)

		t.Run(testname, func(t *testing.T) {
			ch := sleepSort(tt.value)
			for gotValue := range ch {
				got = append(got, gotValue)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error in test test: %d", idx+1)
			}
		})
	}
}
