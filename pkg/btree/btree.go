package btree

/*
	I dont think Nk is being properly updated to replace and just check len keys
*/

type Node struct {
	Nk       int
	Keys     []string
	Children []*Node
	IsLeaf   bool
	mindeg   int
}

type BTree struct {
	Root   *Node
	Count  int
	Height int
	MinDeg int
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
	return &BTree{
		Root:   nil,
		Count:  0,
		Height: 0,
		MinDeg: deg,
	}
}

// Recursive Binary search returns node and the key position.
// If key does not exist return error. Keeps track of the parent
// for deletion cases.
// TODO: Write test and return error
func (b *BTree) search(n *Node, key string, pn *Node) (*Node, int, *Node) {
	low, high := 0, n.Nk

	for low < high {
		mid := low + (high-low)/2
		if key == n.Keys[mid] {
			return n, mid, pn
		} else if key < n.Keys[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low < n.Nk && key == n.Keys[low] {
		return n, low, pn
	}

	if n.IsLeaf {
		//key doesnt exist so maybe return error type
		return nil, -1, pn
	}

	return b.search(n.Children[low], key, n)
}

// Similar to search, returns node and an insert position. Will
// split target nodes Children. Errors if detects duplicates
// TODO: Refresh on split reasoning.
func (b *BTree) searchInsertPos(n *Node, key string) (*Node, int) {
	low, high := 0, n.Nk

	for low < high {
		mid := low + (high-low)/2
		if key == n.Keys[mid] {
			// Duplicate key, error?
			return nil, -1
		} else if key < n.Keys[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low < n.Nk && key == n.Keys[low] {
		// Duplicate
		return nil, -1
	}

	if n.IsLeaf {
		// return and insert in out func
		return n, low
	}

	// TODO: Refresh why split
	if n.Children[low].Nk == 2*b.MinDeg-1 {
		b.splitChild(n, low)
		if key > n.Keys[low] {
			low++
		}
	}

	return b.searchInsertPos(n.Children[low], key)
}

// TODO: Insert tuple k string, v record
func (b *BTree) Insert(key string) {
	if b.Root == nil {
		b.Root = createNode(true, key)
		b.Height++
	} else {	
		if b.Root.Nk > b.MinDeg*2-1 {
			b.Root = b.splitRoot()
		}

		b.insertNotFull(b.Root, key)
	}
	
	b.Count++
}

func (b *BTree) insertNotFull(root *Node, k string) {
	n, i := b.searchInsertPos(root, k)

	var left []string
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
}

// case: reach leaf node, reach internal contains target,
// reach internal no target
// If Children are lacking min node
// 3. if root is deleted? and has Children 1 vs more
func (b *BTree) Delete(n *Node, key string, pn *Node) {
	tNode, kpos, pNode := b.search(n, key, pn)
	//TODO: Error no key

	if tNode.IsLeaf {
		tNode.Keys = append(tNode.Keys[:kpos], tNode.Keys[kpos+1:]...)
		tNode.Nk--
		if tNode.Nk < b.MinDeg-1 && pNode != nil {
			//redistro might still result in less keys than desired
			if ok := pn.redistributeChild(tNode); !ok || tNode.Nk < b.MinDeg {
				b.findMerge(tNode.Keys[0], pn)
			}
		}
		return
	}

	if len(tNode.Children[kpos].Keys) >= n.mindeg {

		predk := b.getPred(tNode, kpos)
		b.Delete(tNode.Children[kpos], predk, pNode)
		tNode.Keys[kpos] = predk
	} else if len(tNode.Children[kpos+1].Keys) >= n.mindeg {

		succk := b.getSucc(tNode, kpos)
		b.Delete(tNode.Children[kpos+1], succk, nil)
		tNode.Keys[kpos] = succk
	} else {

		b.findMerge(tNode.Keys[0], pn)
		if b.Root == tNode && b.Root.Nk == 0 {
			b.Root = tNode.Children[0]
		}

		b.Delete(tNode.Children[kpos], key, nil)
	}
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
	nNode.Nk = b.MinDeg - 1
	nNode.Keys = fullNode.Keys[b.MinDeg:]

	fullNode.Nk = b.MinDeg - 1
	n.insertChild(nNode, i+1)
	n.insertKey(i, fullNode.Keys[b.MinDeg - 1])
	fullNode.Keys = fullNode.Keys[:b.MinDeg-1]
	
	if !fullNode.IsLeaf {
		nNode.Children = fullNode.Children[b.MinDeg:]
		fullNode.Children = fullNode.Children[:b.MinDeg]
	}
}

func (b *BTree) mergeChild(n *Node, i int) {
	//What is we are at last child
	child, sibling := n.Children[i], n.Children[i+1]

	for i, key := range sibling.Keys {
		child.Keys[(b.MinDeg-1)+i] = key
	}

	if !child.IsLeaf {
		for i, child := range sibling.Children {
			child.Children[b.MinDeg+i] = child
		}
	}

}

func (n *Node) redistributeChild(child *Node) bool {
	var childpos int
	var lsib, rsib *Node

	for _, key := range n.Keys {
		if child.Keys[0] > key {
			childpos++
		}
	}

	if childpos > 0 {
		lsib = n.Children[childpos-1]
	}

	if childpos < len(n.Children)-1 {
		rsib = n.Children[childpos+1]
	}

	//take sibling key make parent separator / place parent separater within child
	// Adjust Children?
	// Q How do you insert Keys
	if lsib != nil && len(lsib.Keys) > n.mindeg-1 {
		l := len(lsib.Keys) - 1
		pk := []string{n.Keys[childpos-1]}
		lk := lsib.Keys[l]

		lsib.Keys = lsib.Keys[:l]
		n.Keys[childpos-1] = lk
		//append to beginning
		child.Keys = append(pk, child.Keys...)

		//siblings child given to child
		if !lsib.IsLeaf {
			n.Children[childpos-1] = lsib.Children[l]
		}
		return true
	} else if rsib != nil && len(rsib.Keys) > n.mindeg-1 {
		rk := rsib.Keys[0]
		pk := n.Keys[childpos+1]

		rsib.Keys = rsib.Keys[1:]
		n.Keys[childpos+1] = rk
		child.Keys = append(child.Keys, pk)

		if !rsib.IsLeaf {
			n.Children[childpos+1] = rsib.Children[0]
		}
		return true
	}

	return false
}

func (b *BTree) findMerge(key string, pn *Node) {
	var cPos int
	for k, v := range pn.Children {
		if v.Keys[0] == key {
			cPos = k
		}
	}
	b.mergeChild(pn, cPos)
}

func (n *Node) insertChild(child *Node, i int) {
	if len(n.Children) == i {
		n.Children = append(n.Children, child)
		return
	}

	left := n.Children[:i]
	left = append(left, child)
	left = append(left, n.Children[i:]...)
	n.Children = left
}

//
func (n *Node) insertKey(pos int, key string){
	//might just need one of these
	if len(n.Keys) == 0 || pos == len(n.Keys) {
		n.Keys = append(n.Keys, key)
		n.Nk++
		return
	}
	left := n.Keys[:pos]
	left = append(left, key)
	left = append(left, n.Keys[pos:]...)
	n.Keys = left
	n.Nk++
}
