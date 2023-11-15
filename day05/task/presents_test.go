package task

import (
	cheap "day05/common/heap"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGetNCoolestPresents(t *testing.T) {

	t1 := []cheap.Present{
		{
			Value: 5,
			Size:  1,
		},
		{
			Value: 4,
			Size:  5,
		},
		{
			Value: 3,
			Size:  1,
		},
		{
			Value: 5,
			Size:  2,
		},
	}

	var tests = []struct {
		got         []cheap.Present
		poppedCount int
		want        []cheap.Present
	}{
		{t1, 4, []cheap.Present{{5, 1}, {5, 2}, {4, 5}, {3, 1}}},
		{t1, 2, []cheap.Present{{5, 1}, {5, 2}}},
	}

	for id, tt := range tests {

		testname := fmt.Sprintf("getNcollestPresents_t%d", id+1)
		t.Run(testname, func(t *testing.T) {
			ans, err := getNCoolestPresents(tt.got, tt.poppedCount)
			if !reflect.DeepEqual(tt.want, ans) {
				t.Errorf("want %v, got %v", tt.want, ans)
			}
			if err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("getNcollestPresents_error", func(t *testing.T) {
		var test = []struct {
			got         []cheap.Present
			poppedCount int
			want        []cheap.Present
		}{
			{t1, 0, []cheap.Present{}},
		}

		ans, err := getNCoolestPresents(test[0].got, test[0].poppedCount)

		if len(ans) != 0 && !strings.Contains(err.Error(), "invalid size") {
			t.Errorf("invalid test")
		}
	})
}
