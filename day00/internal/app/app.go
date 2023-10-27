package app

import (
	"Day00/internal/statistic"
	"flag"
	"fmt"

	"github.com/montanaflynn/stats"
)

type Flags struct {
	MeanOption   bool
	MedianOption bool
	ModeOption   bool
	SDOption     bool
}

func Run() {
	var (
		inputFlags Flags
		data       statistic.StatParameters
	)

	flag.BoolVar(&inputFlags.MeanOption, "mean", false, "output mean")
	flag.BoolVar(&inputFlags.MedianOption, "median", false, "output median")
	flag.BoolVar(&inputFlags.ModeOption, "mode", false, "output mode")
	flag.BoolVar(&inputFlags.SDOption, "sd", false, "output sd")
	flag.Parse()

	data.Init()

	if !inputFlags.MeanOption && !inputFlags.MedianOption && !inputFlags.ModeOption && !inputFlags.SDOption {
		mean, err := stats.Mean(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Mean: %.2f\n", mean)

		median, err := stats.Median(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Median: %.2f\n", median)

		mode, err := stats.Mode(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Mode:", mode)

		sd, err := stats.StandardDeviation(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("SD: %.2f\n", sd)
	}

	if inputFlags.MeanOption {
		mean, err := stats.Mean(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Mean: %.2f\n", mean)
	}

	if inputFlags.MedianOption {
		median, err := stats.Median(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Median: %.2f\n", median)
	}

	if inputFlags.ModeOption {
		mode, err := stats.Mode(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Mode:", mode)
	}

	if inputFlags.SDOption {
		sd, err := stats.StandardDeviation(data.Array)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("SD: %.2f\n", sd)
	}
}
