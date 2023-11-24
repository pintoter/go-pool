bench:
		go test -bench=. min_coins.go min_coins_b_test.go -cpuprofile=cpu.out -memprofile=mem.out
		go tool pprof cpu.out
		top10 > top10.txt