package btree_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/btree"
	"github.com/stretchr/testify/assert"
)

/*
TODO: Test
	[X]Insert and split child
		[X] Duplicate error
	[X] Search
		[X] Key doesnt exist
	[X] Split Child(not root)
		[X] Test split grandchildren(height 3)
	[] Delete
		[] Redistro case left vs right
			Redistro leaf vs internal
		[] Merge after redistro case


*/

func TestSearchTree(t *testing.T) {
	tests := []struct {
		minDeg    int
		init      []string
		xh int
		search []string
	}{
		{
			5,
			[]string{"101", "20", "457", "500", "100", "30", "45", "50", "60", "66","1000", "2000", "3000", "4000", "5000", "6000"},
			2,
			[]string{"100", "1000", "101", "45","457", "50", "500", "5000", "60", "6000", "66"},
		},
		{
			2,
			[]string{"101", "20", "457", "500","1000", "2000", "3000", "4000", "5000", "6000", "7",
			"70", "100", "30", "45", "50", "60", "66", "67", "68"},
			3,
			[]string{"100", "1000", "101", "20", "2000", "30", "3000", "4000", "45",
			"457", "50", "500", "5000", "60", "6000", "66", "67", "68", "7", "70"},
		},
	}

	for _, tt := range tests {
		tbt := btree.CreateBTree(tt.minDeg)

		for _, v := range tt.init {
			err := tbt.Insert(v)
			assert.Equal(t, nil, err)
		}

		assert.Equal(t, tt.xh, tbt.Height)

		for _, v := range tt.search {
			n, i, _, err := tbt.Search(tbt.Root, v, nil)
			assert.Equal(t, nil, err)
			assert.Equal(t, n.Keys[i], v)
		}

		_,_,_, err := tbt.Search(tbt.Root, "900", nil)
		assert.EqualError(t, err, btree.ErrKeyDoesntExist)
		//check for existing key in wrong side of tree
		_,_,_, err = tbt.Search(tbt.Root.Children[0], "66", tbt.Root)
		assert.EqualError(t, err, btree.ErrKeyDoesntExist)


	}
}

func TestBTreeInsert(t *testing.T) {
	tbt := btree.CreateBTree(5)

	var err error
	for _, v := range []string{"100", "50", "45", "60", "30"} {
		err = tbt.Insert(v)
	}

	assert.Nil(t, err)

	for i, tc := range []string{"100", "30", "45", "50", "60"} {
		if tbt.Root.Keys[i] != tc {
			t.Fatalf("Key ordering not correct key [%d] = %s, not %s, ", i, tbt.Root.Keys[i], tc)
		}
	}

	for _, v := range []string{"101", "20", "457", "500"} {
		err = tbt.Insert(v)
	}

	assert.Nil(t, err)

	for i, tc := range []string{"100", "101", "20", "30", "45", "457", "50", "500", "60"} {
		if tbt.Root.Keys[i] != tc {
			t.Fatalf("Key ordering not correct key [%d] = %s, not %s, ", i, tbt.Root.Keys[i], tc)
		}
	}

	err = tbt.Insert("100")
	assert.EqualError(t, err, btree.ErrDuplicateKey)
	err = tbt.Insert("20")
	assert.EqualError(t, err, btree.ErrDuplicateKey)
	err = tbt.Insert("500")
	assert.EqualError(t, err, btree.ErrDuplicateKey)

}

func TestBTreeSplitRoot(t *testing.T) {
	tests := []struct {
		minDeg    int
		init      []string
		xh        int
		insert    []string
		xkeyOrder []string
	}{
		{
			5,
			[]string{"101", "20", "457", "500", "100", "30", "45", "50", "60", "66"},
			2,
			[]string{"1000"},
			[]string{"100", "1000", "101", "20", "30", "45", "457", "50", "500", "60", "66"},
		},
		{
			5,
			[]string{"101", "20", "457", "500", "100", "30", "45", "50", "60", "66"},
			2,
			[]string{"1000", "2000", "3000", "4000", "5000", "6000"},
			[]string{"100", "1000", "101", "20", "2000", "30", "3000", "4000", "45",
				"457", "50", "500", "5000", "60", "6000", "66"},
		},
		{
			2,
			[]string{"101", "20", "457", "500"},
			3,
			[]string{"1000", "2000", "3000", "4000", "5000", "6000", "7",
				"70", "100", "30", "45", "50", "60", "66", "67", "68"},
			[]string{"100", "1000", "101", "20", "2000", "30", "3000", "4000", "45",
				"457", "50", "500", "5000", "60", "6000", "66", "67", "68", "7", "70"},
		},
	}

	// TODO: Finishing out test case
	for _, tt := range tests {
		tbt := btree.CreateBTree(tt.minDeg)

		for _, v := range tt.init {
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
			t.Errorf("expected root to have 1 key; got [%d]", tbt.Root.Nk)
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
	tests := []struct {
		minDeg int
		init   []string
		xh     int
		numC   int
		order  []string
	}{
		{
			3,
			[]string{"101", "20", "457", "500", "100", "30", "45", "102", "1000", "103", "104", "105"},
			2,
			3,
			[]string{"100", "1000", "101", "102", "103", "104", "105", "20", "30", "45", "457", "500"},
		},
		{
			2,
			[]string{"101", "20", "457", "500", "100", "50", "51", "53"},
			2,
			3,
			[]string{"100", "101", "20", "457", "50", "500", "51", "53"},
		},
	}

	for _, tt := range tests {
		tbt := btree.CreateBTree(tt.minDeg)

		for _, v := range tt.init {
			tbt.Insert(v)
		}

		if tbt.Height != tt.xh {
			t.Errorf("expected height %d; got [%d], #keys:[%d]", tt.xh, tbt.Height, tbt.Root.Nk)
		}

		if len(tbt.Root.Children) != tt.numC {
			t.Errorf("expected number of children %d; got [%d]", tt.numC, len(tbt.Root.Children))
		}

		tOrder := inOrderTrav(tbt.Root)

		for i, v := range tt.order {
			if tOrder[i] != v {
				t.Errorf("keys out of order. Expected [%s] at idx [%d]; got [%s]", v, i, tOrder[i])
				t.Log(tOrder)
				return
			}
		}

	}
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
