package statistic

import (
	"Day00/internal/model"
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MAX = 100000
	MIN = -100000
)

type StatParameters struct {
	Array []float64
}

func (sp *StatParameters) Init() {
	in := bufio.NewReader(os.Stdin)

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		line = strings.TrimSpace(line)

		inpNumber, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(model.NotDigitInput)
			continue
		}

		if inpNumber > MAX || inpNumber < MIN {
			fmt.Println(model.DigitOverLimit)
			continue
		}

		sp.Array = append(sp.Array, float64(inpNumber))
	}

	sort.Float64s(sp.Array)
}
