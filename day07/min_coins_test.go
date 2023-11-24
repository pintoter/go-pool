package main

import (
	"fmt"
	"reflect"
	"testing"
)

type test struct {
	coins []int
	value int
	want  []int
}

var tests = []test{
	{
		value: 13,
		coins: []int{},
		want:  []int{},
	},
	{
		value: 13,
		coins: []int{1, 5, 10},
		want:  []int{1, 1, 1, 10},
	},
	{
		value: 15,
		coins: []int{1, 5, 10},
		want:  []int{5, 10},
	},
	{
		value: 18,
		coins: []int{1, 6, 10},
		want:  []int{6, 6, 6},
	},
}

func TestMinCoins(t *testing.T) {
	for idx, tc := range tests {
		testname := fmt.Sprintf("MinCoinsTest#%d", idx+1)
		t.Run(testname, func(t *testing.T) {
			got := MinCoins(tc.value, tc.coins)
			if !reflect.DeepEqual(tc.want, got) {
				t.Error()
			}
		})
	}
}

func TestMinCoins2(t *testing.T) {
	for idx, tc := range tests {
		testname := fmt.Sprintf("MinCoins2Test#%d", idx+1)
		t.Run(testname, func(t *testing.T) {
			got := MinCoins2(tc.value, tc.coins)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("got equal %v", got)
			}
		})
	}
}
