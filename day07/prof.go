/*
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"testing"
	"time"

	"github.com/pkg/profile"
)

func BenchmarkMinCoins(b *testing.B) {
	prof := profile.Start(profile.CPUProfile)

	coins := append([]int{1}, getRandomData2(1000)...)
	val := 12345

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinCoins2(val, coins)
	}
	prof.Stop()

	output, err := exec.Command("go tool pprof", "-top", "-unit", "10", "./cpu.pprof").Output()
	// cmd.Stdin = os.Stdin

	// file, err := os.Create("top10.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd.Stdout = file
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// output, err := cmd.Output()

	// err = cmd.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(output))
	// cmd.Close()
}

func getRandomData2(cap int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, cap)
	for i := 0; i < cap; i++ {
		res[i] = rand.Int()
	}
	return res
}
*/