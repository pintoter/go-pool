package main

import (
	"math/rand"
	"testing"
	"time"
)

var coins = append([]int{1}, getRandomData(1000)...)

func BenchmarkMinCoins(b *testing.B) {
	val := 12345

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = MinCoins(val, coins)
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	val := 12345

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = MinCoins2(val, coins)
	}
}

func getRandomData(cap int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, cap)
	for i := 0; i < cap; i++ {
		res[i] = rand.Int()
	}
	return res
}
