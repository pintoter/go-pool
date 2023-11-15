package task

import (
	"day05/common"
	"fmt"
	"reflect"
	"testing"
)

func TestUnrollGirland(t *testing.T) {
	/*
					____1____
				 /		     \
				0			      1
			 / \		    /	   \
		  1	  1      1		  0
		 /\		/\    /\		  /\
		1  0 0  1  0  0    1  1
	*/
	t1 := common.Create(true)
	{
		t1.Left = common.Create(false)
		t1.Left.Left = common.Create(true)
		t1.Left.Left.Left = common.Create(true)
		t1.Left.Left.Right = common.Create(false)
		t1.Left.Right = common.Create(true)
		t1.Left.Right.Left = common.Create(false)
		t1.Left.Right.Right = common.Create(true)
		t1.Right = common.Create(true)
		t1.Right.Left = common.Create(true)
		t1.Right.Left.Left = common.Create(false)
		t1.Right.Left.Right = common.Create(false)
		t1.Right.Right = common.Create(false)
		t1.Right.Right.Left = common.Create(true)
		t1.Right.Right.Right = common.Create(true)
	}

	var tests = []struct {
		t    *common.TreeNode
		want []bool
	}{
		{t1, []bool{true, false, true, false, true, true, true, true, false, false, true, false, false, true, true}},
	}

	for id, tt := range tests {

		testname := fmt.Sprintf("unrollGirland_t%d", id+1)
		t.Run(testname, func(t *testing.T) {
			ans := unrollGirliand(tt.t)
			if !reflect.DeepEqual(tt.want, ans) {
				t.Errorf("want %v, got %v", tt.want, ans)
			}
		})
	}
}

/*

want [true false true false true true true true false false true false false true true], 
got  [true false true false true true true false true true false false false true true]

*/