package btree

import "errors"


const (
	ErrDuplicateKey = "duplicate key found"
	ErrKeyDoesntExist = "key does not exist"
)

type Node struct {
	Nk       int
	Keys     []string
	Children []*Node
	IsLeaf   bool
}

type BTree struct {
	Root   *Node
	Count  int // num of nodes?
	Height int
	MinDeg int // defines lower and upper bound
}

func createNode(leaf bool, keys ...string) *Node {
	return &Node{
		Nk:       len(keys),
		Keys:     keys,
		Children: []*Node{},
		IsLeaf:   leaf,
	}
}

func CreateBTree(deg int) *BTree {
	//min
	return &BTree{
		Root:   nil,
		Count:  0,
		Height: 0,
		MinDeg: deg,
	}
}

// Recursive Binary Search returns node and the key position.
// If key does not exist return error. Keeps track of the parent
// for deletion cases.
// Returns Node, Node position, Parent Node and error
func (b *BTree) Search(n *Node, key string, pn *Node) (*Node, int, *Node, error) {
	low, high := 0, n.Nk

	for low < high {
		mid := low + (high-low)/2
		if key == n.Keys[mid] {
			return n, mid, pn, nil
		} else if key < n.Keys[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low < n.Nk && key == n.Keys[low] {
		return n, low, pn, nil
	}

	if n.IsLeaf {
		//TODO: Consider better error building or handling
		return nil, -1, pn, errors.New(ErrKeyDoesntExist)
	}

	return b.Search(n.Children[low], key, n)
}

// Similar to Search, returns node and an insert position. 
// Will split target nodes children and rebalance tree as
// the Search descends.
func (b *BTree) searchInsertPos(n *Node, key string) (*Node, int, error) {
	low, high := 0, n.Nk

	for low < high {
		mid := low + (high-low)/2
		if key == n.Keys[mid] {
			return nil, -1, errors.New(ErrDuplicateKey)
		} else if key < n.Keys[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low < n.Nk && key == n.Keys[low] {
		return nil, -1, errors.New(ErrDuplicateKey)
	}

	if n.IsLeaf {
		return n, low, nil
	}

	if n.Children[low].Nk == 2*b.MinDeg-1 {
		// Split before any insert also potential rebalance
		// tree as descend.
		b.splitChild(n, low)
		if key > n.Keys[low] {
			low++
		}
	}

	return b.searchInsertPos(n.Children[low], key)
}

// TODO: Insert tuple k string, v record
func (b *BTree) Insert(key string) error {
	var err error
	if b.Root == nil {
		b.Root = createNode(true, key)
		b.Height++
	} else {	
		if b.Root.Nk == b.MinDeg*2-1 {
			b.Root = b.splitRoot()
		}

		err = b.insertNotFull(b.Root, key)
	}
	
	b.Count++
	return err
}

func (b *BTree) insertNotFull(root *Node, k string) error {
	n, i, err := b.searchInsertPos(root, k)
	if err != nil {
		return err
	}

	var left []string
	//To insert before or after end
	if i == n.Nk {
		left = append(n.Keys, k)

	} else if k < n.Keys[i] {
		left = make([]string, len(n.Keys[:i]))
		copy(left, n.Keys[:i])
		left = append(left, k)
		left = append(left, n.Keys[i:]...)

	} else {
		left = make([]string, len(n.Keys[:i+1]))
		copy(left, n.Keys[:i+1])
		left = append(left, k)
		left = append(left, n.Keys[i+1:]...)
	}

	n.Keys = left
	n.Nk++
	// disk write?
	return nil
}

// I suspect that this function has all the moving parts
// but isnt truly doing its job. Requires thorough testing
// Seems like pn purpose is to ensure that search has a parent
// node to return for case when it doesnt recurse.
func (b *BTree) Delete(n *Node, key string, pn *Node) error {
	var err error
	kpos := 0
	for kpos < n.Nk && key > n.Keys[kpos] {
		kpos++
	}

	//  Key is in current node
	if kpos < n.Nk && n.Keys[kpos] == key {
		if n.IsLeaf {
			n.Keys = append(n.Keys[:kpos], n.Keys[kpos+1:]...)
			n.Nk--
			b.Count--
			return nil
		}

		// Internal node handling
		left, right := n.Children[kpos], n.Children[kpos+1]
		if left.Nk >= b.MinDeg{
			predk := b.getPred(n, kpos)
			err = b.Delete(left, predk, n)
			if err !=nil {
				return err
			}

			n.Keys[kpos] = predk
			return nil
		}

		
		if right.Nk >= b.MinDeg {
		 	succk := b.getSucc(n, kpos)
			err := b.Delete(right , succk, n)
			if err != nil {
				return err
			}

			n.Keys[kpos] = succk
			return nil
		}

		// Both children have a deficit
		b.mergeChild(n, kpos)
		if b.Root == n && n.Nk == 0 {
			b.Root = n.Children[0]
			n = b.Root
		}

		return b.Delete(n.Children[kpos], key, n)
	}

	if n.IsLeaf {
		return errors.New(ErrKeyDoesntExist)
	}

	child := n.Children[kpos]
	
	// correct child
	if child.Nk == b.MinDeg-1 {
		ok := n.redistributeChild(child, b.MinDeg)
		if !ok {
			if kpos == n.Nk{
				b.mergeChild(n, kpos-1)
				kpos--
			} else{
				b.mergeChild(n, kpos)
			}

			child = n.Children[kpos]

			if n == b.Root && n.Nk == 0 {
				b.Root = child
			}
		}
	}


	return b.Delete(child, key, n)
}

func (b *BTree) getPred(n *Node, idx int) string {
	curr := n.Children[idx]

	for !curr.IsLeaf {
		curr = curr.Children[curr.Nk]
	}

	return curr.Keys[curr.Nk-1]
}

func (b *BTree) getSucc(n *Node, idx int) string {
	curr := n.Children[idx+1]

	for !curr.IsLeaf {
		curr = curr.Children[0]
	}

	return curr.Keys[0]
}

func (b *BTree) splitRoot() *Node {
	nRoot := createNode(false)
	nRoot.Children = append(nRoot.Children, b.Root)
	b.Root = nRoot
	b.splitChild(nRoot, 0)
	b.Height++
	return nRoot
}

func (b *BTree) splitChild(n *Node, i int) {
	fullNode := n.Children[i]
	nNode := createNode(fullNode.IsLeaf)
	t := b.MinDeg

	promoted := fullNode.Keys[t-1]
	
	nNode.Keys = make([]string, len(fullNode.Keys[t:]))
	copy(nNode.Keys, fullNode.Keys[t:])
	nNode.Nk = len(nNode.Keys)

	fullNode.Keys = fullNode.Keys[:t-1]
	fullNode.Nk = len(fullNode.Keys)

	fullNode.Keys = fullNode.Keys[:b.MinDeg-1]
	
	if !fullNode.IsLeaf {
		nNode.Children = make([]*Node, len(fullNode.Children[:t]))
		copy(nNode.Children, fullNode.Children[t:])

		fullNode.Children = fullNode.Children[:t]
	}

	n.insertChild(nNode, i+1)
	n.insertKey(i, promoted)
}

func (b *BTree) mergeChild(n *Node, i int) {
	
	if i == len(n.Children)-1 {
		i--
	}

	left, right := n.Children[i], n.Children[i+1]
	parentKey := n.Keys[i]


	left.Keys = append(left.Keys,parentKey)
	left.Keys = append(left.Keys,right.Keys...)

	if !left.IsLeaf{
		left.Children = append(left.Children, right.Children...)
	}

	n.Keys = append(n.Keys[:i], n.Keys[i+1:]...)
	n.Nk--

	n.Children = append(n.Children[:i+1],n.Children[i+2:]...)
	left.Nk = len(left.Keys)
	b.Count--

	if len(n.Children) != len(n.Keys)+1 {
    	panic("B-tree invariant violated")
	}
}

func (n *Node) redistributeChild(child *Node, mindeg int) bool {
	var childpos int
	var lsib, rsib *Node

	if len(child.Keys) != 0 {
		for _, key := range n.Keys {
			if child.Keys[0] > key {
				childpos++
			}
		}
	}
	
	if childpos > 0 {
		lsib = n.Children[childpos-1]
	}

	if childpos < len(n.Children)-1 {
		rsib = n.Children[childpos+1]
	}

	//take sibling key make parent separator / place parent separater within child
	if lsib != nil && len(lsib.Keys) >= mindeg {
		l := len(lsib.Keys) - 1
		pk := []string{n.Keys[childpos-1]}
		lk := lsib.Keys[l]

		// Movinf sibling key into parent to make it a separator
		lsib.Keys = lsib.Keys[:l]
		n.Keys[childpos-1] = lk
		child.Keys = append(pk, child.Keys...)

		//Adjust children
		if !lsib.IsLeaf {
			child.Children = append([]*Node{lsib.Children[l]}, child.Children...)
			lsib.Children = lsib.Children[:l]
		}
		
		lsib.Nk--
		child.Nk++

		return true
	} else if rsib != nil && len(rsib.Keys) >= mindeg {
		rk := rsib.Keys[0]
		pk := n.Keys[childpos]

		rsib.Keys = rsib.Keys[1:]
		n.Keys[childpos] = rk
		child.Keys = append(child.Keys, pk)

		if !rsib.IsLeaf {
			child.Children = append(child.Children, rsib.Children[0])
			rsib.Children = rsib.Children[1:]
		}

		rsib.Nk--
		child.Nk++

		return true
	}

	return false
}

func (n *Node) insertChild(child *Node, i int) {
	if len(n.Children) == i {
		n.Children = append(n.Children, child)
		return
	}

	left := make([]*Node, len(n.Children[:i]))
	copy(left,n.Children[:i])
	left = append(left, child)
	left = append(left, n.Children[i:]...)
	n.Children = left
}

//
func (n *Node) insertKey(pos int, key string){
	if len(n.Keys) == 0 || pos > n.Nk-1 {
		n.Keys = append(n.Keys, key)

		n.Nk = len(n.Keys)
		return
	}

	var left []string
	if key < n.Keys[pos]{
		left = make([]string, len(n.Keys[:pos]))
		copy(left, n.Keys[:pos])
		left = append(left, key)
		left = append(left, n.Keys[pos:]...)
	} else {
		left = make([]string, len(n.Keys[:pos+1]))
		copy(left,n.Keys[:pos+1])
		left = append(left, key)
		left = append(left, n.Keys[pos+1:]...)
	}
	
	n.Keys = left
	n.Nk++
}
