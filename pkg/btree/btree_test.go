package btree_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/btree"
)

/* 
TODO: Test
	Insert and split child
	Search
	Split Child(not root)
	Delete
		Redistro case left vs right
			Redistro leaf vs internal
		Merge after redistro case


*/

func TestBTreeInsert(t *testing.T) {
	tbt := btree.CreateBTree(5)

	for _, v := range []string{"100", "50", "45", "60", "30"} {
		tbt.Insert(v)
	}

	for i, tc := range []string{"100", "30", "45", "50", "60"} {
		if tbt.Root.Keys[i] != tc {
			t.Fatalf("Key ordering not correct key [%d] = %s, not %s, ", i, tbt.Root.Keys[i], tc)
		}
	}

	for _, v := range []string{"101", "20", "457", "500"} {
		tbt.Insert(v)
	}

	for i, tc := range []string{"100", "101", "20", "30", "45", "457", "50", "500", "60"} {
		if tbt.Root.Keys[i] != tc {
			t.Fatalf("Key ordering not correct key [%d] = %s, not %s, ", i, tbt.Root.Keys[i], tc)
		}
	}

}

func TestBTreeSplitRoot(t *testing.T) {
	tests := []struct{
		minDeg int
		noSplit []string
		xh int
		insert []string
		xkeyOrder []string
	}{
		{
			5, 
			[]string{"101", "20", "457", "500", "100", "30", "45", "50", "60", "66"},
			2,
			[]string{"1000"},
			[]string{"100", "1000", "101", "20", "30", "45", "457", "50","500", "60", "66"},
		},
		{
			5, 
			[]string{"101", "20", "457", "500", "100", "30", "45", "50", "60", "66"},
			2,
			[]string{"1000", "2000", "3000", "4000", "5000", "6000"},
			[]string{"100", "1000", "101", "20","2000", "30", "3000", "4000","45", 
			"457", "50","500","5000", "60", "6000", "66"},
		},
	}

	// TODO: Finishing out test case
	for _, tt := range tests {	
		tbt := btree.CreateBTree(tt.minDeg)

		for _, v := range tt.noSplit{
			tbt.Insert(v)
		}

		if tbt.Height != 1 {
			t.Errorf("expected height 1; got [%d], #keys:[%d]", tbt.Height, tbt.Root.Nk)
		}

		for _, v := range tt.insert {
			tbt.Insert(v)
		}

		if tbt.Height != tt.xh {
			t.Log(tbt.Root.Keys)
			t.Log(tbt.Root.Children)
			t.Errorf("expected height [%d]; got [%d], #keys:[%d]", tt.xh, tbt.Height, tbt.Root.Nk)
		}

		if tbt.Root.Nk != 1 {
			t.Log(tbt.Root.Keys)
			t.Errorf("expected root to have 1 key; got [%d]",tbt.Root.Nk)
		}

		order := inOrderTrav(tbt.Root)

		for i, v := range tt.xkeyOrder {
			if order[i] != v {
				t.Errorf("keys out of order. Expected [%s] at idx [%d]; got [%s]", v, i, order[i])
				t.Log(order)
				return
			}
		}
	}
}

func TestBTreeSplitChild(t *testing.T) {
	
}

func inOrderTrav(node *btree.Node) []string {
	if len(node.Children) == 0 {
		return node.Keys
	}

	var res []string

	for i, v := range node.Children {
		left := inOrderTrav(v)
		res = append(res, left...)
		if i != len(node.Keys) {
			res = append(res, node.Keys[i])
		}
	}

	return res
}
// func display(t *testing.T, node *btree.Node) {
// 	if node == nil {
// 		return
// 	}

// 	var i int
// 	for i = 0; i < node.Nk; i++ {
// 		if !node.IsLeaf {
// 			display(t, node.Children[i])
// 		}
// 		t.Log(node.Keys[i])
// 	}
// }
