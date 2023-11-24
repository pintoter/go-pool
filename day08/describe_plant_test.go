package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDescribePlant(t *testing.T) {
	var tests = []struct {
		value interface{}
		want  string
	}{
		{
			AnotherUnknownPlant{
				FlowerColor: 10,
				LeafType:    "lanceolate",
				Height:      100,
			},
			"FlowerColor:10\nLeafType:lanceolate\nHeight(unit=inches):100\n",
		},
		{
			UnknownPlant{
				FlowerType: "peony",
				LeafType:   "the leaves are compound (made up of two or more discrete leaflets",
				Color:      14501017,
			},
			"FlowerType:peony\nLeafType:the leaves are compound (made up of two or more discrete leaflets\nColor(color_scheme=rgb):14501017\n",
		},
		{
			[]int{1, 2, 3},
			"unsupported type: []int\n",
		},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("getElement_test#%d", idx+1)
		t.Run(testname, func(t *testing.T) {
			old := os.Stdout // keep backup of the real stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			describePlant(tt.value)

			outC := make(chan string)
			// copy the output in a separate goroutine so printing can't block indefinitely
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()

			w.Close()
			os.Stdout = old // restoring the real stdout
			res := <-outC

			if res != tt.want {
				t.Errorf("got [%s] = want [%s]\n", res, tt.want)
			}
		})
	}

}
