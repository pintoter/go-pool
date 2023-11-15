package task

import (
	"day05/common/tree"
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
									_______1_______
								 /			     	  	\
							   0			      	     1
							 /	 \		    		 /	    \
					    1	    1     		  1		      0
					 	 /\		  /\    		 /\		       /\
						/	 \ 	 /  \			  / 	\		    /	 \
					 1    0  0   1 	   0     0     1     1
		      /\   /\  /\  /\   /\     /\   /\     /\
		     0  0 0 0  0 0 0 0 1  1   1  1 0 0    1  1
	*/
	t2 := tree.Create(true)
	{
		t2.Left = tree.Create(false)
		t2.Left.Left = tree.Create(true)
		t2.Left.Left.Left = tree.Create(true)
		t2.Left.Left.Left.Left = tree.Create(false)
		t2.Left.Left.Left.Right = tree.Create(false)
		t2.Left.Left.Right = tree.Create(false)
		t2.Left.Left.Right.Left = tree.Create(false)
		t2.Left.Left.Right.Right = tree.Create(false)
		t2.Left.Right = tree.Create(true)
		t2.Left.Right.Left = tree.Create(false)
		t2.Left.Right.Left.Left = tree.Create(false)
		t2.Left.Right.Left.Right = tree.Create(false)
		t2.Left.Right.Right = tree.Create(true)
		t2.Left.Right.Right.Left = tree.Create(false)
		t2.Left.Right.Right.Right = tree.Create(false)
		t2.Right = tree.Create(true)
		t2.Right.Left = tree.Create(true)
		t2.Right.Left.Left = tree.Create(false)
		t2.Right.Left.Left.Left = tree.Create(true)
		t2.Right.Left.Left.Right = tree.Create(true)
		t2.Right.Left.Right = tree.Create(false)
		t2.Right.Left.Right.Left = tree.Create(true)
		t2.Right.Left.Right.Right = tree.Create(true)
		t2.Right.Right = tree.Create(false)
		t2.Right.Right.Left = tree.Create(true)
		t2.Right.Right.Left.Left = tree.Create(false)
		t2.Right.Right.Left.Right = tree.Create(false)
		t2.Right.Right.Right = tree.Create(true)
		t2.Right.Right.Right.Left = tree.Create(true)
		t2.Right.Right.Right.Right = tree.Create(true)
	}

	var tests = []struct {
		t    *tree.Node
		want []bool
	}{
		{t1, []bool{true, false, true, false, true, true, true, true, false, false, true, false, false, true, true}},
		{t2, []bool{true, false, true, false, true, true, true, true, false, false, true, false, false, true,
			true, true, true, false, false, true, true, true, true, false, false, false, false, false, false, false, false}},
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
