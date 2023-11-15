package task

import (
	"day05/common/tree"
	"fmt"
	"testing"
)

func TestAreToysBalanced(t *testing.T) {
	/*
					____1____
				 /		     \
				0			      1
			 / \		    /	   \
		  1	  1      1		  0
		 /\		/\    /\		  /\
		1  0 0  1  0  0    1  1
	*/
	t1 := tree.Create(true)
	{
		t1.Left = tree.Create(false)
		t1.Left.Left = tree.Create(true)
		t1.Left.Left.Left = tree.Create(true)
		t1.Left.Left.Right = tree.Create(false)
		t1.Left.Right = tree.Create(true)
		t1.Left.Right.Left = tree.Create(false)
		t1.Left.Right.Right = tree.Create(true)
		t1.Right = tree.Create(true)
		t1.Right.Left = tree.Create(true)
		t1.Right.Left.Left = tree.Create(false)
		t1.Right.Left.Right = tree.Create(false)
		t1.Right.Right = tree.Create(false)
		t1.Right.Right.Left = tree.Create(true)
		t1.Right.Right.Right = tree.Create(true)
	}
	/*
			____1____
		 /		     \
		0			      1
	*/
	t2 := tree.Create(true)
	{
		t2.Left = tree.Create(false)
		t2.Right = tree.Create(true)
	}

	/*
		___NIL___
	*/
	var t3 *tree.Node

	/*
			____1____
		 /		     \
		1			     NIL
	*/
	t4 := tree.Create(true)
	t4.Left = tree.Create(true)

	/*
			____1____
		 /		     \
		0			     NIL
	*/
	t5 := tree.Create(true)
	t5.Left = tree.Create(false)

	var tests = []struct {
		t    *tree.Node
		want bool
	}{
		{t1, true},
		{t2, false},
		{t3, false},
		{t4, false},
		{t5, true},
	}

	for id, tt := range tests {
		testname := fmt.Sprintf("areToysBalanced_t%d", id+1)
		t.Run(testname, func(t *testing.T) {
			ans := AreToysBalanced(tt.t)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
