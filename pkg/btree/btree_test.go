package btree_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/btree"
)

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
