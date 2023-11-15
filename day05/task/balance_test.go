package task

import (
	"day05/common"
	"fmt"
	"testing"
)

func TestBalance(t *testing.T) {
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
	/*
			____1____
		 /		     \
		0			      1
	*/
	t2 := common.Create(true)
	{
		t2.Left = common.Create(false)
		t2.Right = common.Create(true)
	}

	/*
		___NIL___
	*/
	var t3 *common.TreeNode

	/*
			____1____
		 /		     \
		1			     NIL
	*/
	t4 := common.Create(true)
	t4.Left = common.Create(true)

	/*
			____1____
		 /		     \
		0			     NIL
	*/
	t5 := common.Create(true)
	t5.Left = common.Create(false)

	var tests = []struct {
		t    *common.TreeNode
		want bool
	}{
		{t1, true},
		{t2, false},
		{t3, false},
		{t4, false},
		{t5, true},
	}

	for id, tt := range tests {
		testname := fmt.Sprintf("Tree%d", id)
		t.Run(testname, func(t *testing.T) {
			ans := AreToysBalanced(tt.t)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
