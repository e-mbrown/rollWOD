package btree

import "fmt"

/*
	Reminder
	50
	root 2-m children, 1 - m-1 keys
	nonroot m/2 -m children, m-1 keys

	Working on redistribution, and clearing cases....Then complete deletion then test each part of the b-tree. Then attach data pointers.
*/


type Node struct {
	nk int
	keys []string
	children []*Node
	isLeaf bool
	mindeg int
}

type BTree struct {
	root *Node
	count int
	height int
	minDeg int
}

func createNode(leaf bool) *Node {
	return &Node{
		nk: 0,
		keys: []string{},
		children: []*Node{},
		isLeaf: leaf,
	}
}

func createBTree(deg int) *BTree {
	root := createNode(true)
	return &BTree {
		root: root,
		count: 1,
		height: 0,
		minDeg: deg,
	}
}

// Recursive Binary search returns node and the key position.
// If key does not exist return error. Keeps track of the parent
// for deletion cases.
// TODO: Write test and return error
func (b *BTree) search(n *Node, key string, pn *Node) (*Node, int, *Node){
	low, high := 0, n.nk

	for low < high {
		mid := low + (high-low)/2
		if key == n.keys[mid]{
			return n, mid, pn
		}else if key < n.keys[mid] {
			high = mid
		}else {
			low = mid+1
		}
	}

	if low < n.nk && key == n.keys[low] {
		return n, low, pn
	}

	if n.isLeaf{
		//key doesnt exist so maybe return error type
		return nil, -1, pn
	}

	return b.search(n.children[low], key, n)
}

// Similar to search, returns node and an insert position. Will
// split target nodes children. Errors if detects duplicates
// TODO: Refresh on split reasoning.
func (b *BTree) searchInsertPos(n *Node, key string) (*Node, int){
	low, high := 0, n.nk

	for low < high {
		mid := low + (high-low)/2
		if key == n.keys[mid]{
			// Duplicate key, error?
			return nil,-1
		}else if key < n.keys[mid] {
			high = mid
		}else {
			low = mid+1
		}
	}

	if low < n.nk && key == n.keys[low] {
		// Duplicate
		return nil, -1
	}

	if n.isLeaf{
		// return and insert in out func
		return n, low
	}	
	
	// TODO: Refresh why split
	if n.children[low].nk == 2*b.minDeg-1 {
		b.splitChild(n, low)
		if key > n.keys[low] { low++ }
	}

	return b.searchInsertPos(n.children[low], key)
}

// Rebalances as it descends? I think
func (b *BTree) insert(key string){
	root := b.root
	if root.nk > b.minDeg * 2 {
		root = b.splitRoot()
	}
	
	b.insertNotFull(root, key)
}

func (b *BTree) insertNotFull(node *Node, k string){
	//theoretically node is a leaf and i is
	// key that our key may or may not be larger than
	n, i := b.searchInsertPos(node, k)
	var left []string
	if k < n.keys[i]{
		copy(left,n.keys[:i])
		left = append(left, k)
		left = append(left, n.keys[i:]...)
	} else {
		copy(left, n.keys[:i+1])
		left = append(left, k)
		left = append(left, n.keys[i+1:]...)
	}
	// disk write?
}

//case: reach leaf node, reach internal contains target,
// reach internal no target
// If children are lacking min node
// 3. if root is deleted? and has children 1 vs more
func(b *BTree) delete(n *Node, key string, pn *Node) {	
	tNode, kpos, pNode := b.search(n, key, pn)
	//TODO: Error no key

	if tNode.isLeaf {
		tNode.keys = append(tNode.keys[:kpos], tNode.keys[kpos+1:]...)
		tNode.nk--
		if tNode.nk < b.minDeg-1  && pNode != nil{
			//redistro might still result in less keys than desired
			if ok := pn.redistributeChild(tNode); !ok || tNode.nk < b.minDeg {
				b.findMerge(tNode.keys[0], pn)
			}
		}
		return
	}

	if len(tNode.children[kpos].keys) >= n.mindeg {
		
		predk := b.getPred(tNode, kpos)
		b.delete(tNode.children[kpos], predk, pNode)
		tNode.keys[kpos] = predk
	} else if len(tNode.children[kpos+1].keys) >= n.mindeg{

		succk := b.getSucc(tNode, kpos)
		b.delete(tNode.children[kpos+1],succk, nil)
		tNode.keys[kpos] = succk
	} else {

		b.findMerge(tNode.keys[0], pn)
		if b.root == tNode && b.root.nk == 0 {
			b.root = tNode.children[0]
		}
		
		b.delete(tNode.children[kpos], key, nil)
	}
}
		
func (b *BTree) getPred(n *Node, idx int) string {
	curr := n.children[idx]

	for !curr.isLeaf {
		curr = curr.children[curr.nk]
	}

	return curr.keys[curr.nk-1]
}


func (b *BTree) getSucc(n *Node, idx int) string {
	curr := n.children[idx+1]

	for !curr.isLeaf {
		curr = curr.children[0]
	}

	return curr.keys[0]
}

func (b *BTree) splitRoot() *Node {
	nRoot := createNode(false)
	nRoot.children = append(nRoot.children, b.root)
	b.root = nRoot
	b.splitChild(nRoot, 0)
	return nRoot
}

func (b *BTree) splitChild(n *Node, i int) {
	fullNode := n.children[i]
	nNode := createNode(fullNode.isLeaf)
	nNode.nk = b.minDeg - 1
	//Taking greater keys
	nNode.keys = fullNode.keys[b.minDeg:]
	
	if !fullNode.isLeaf{
		nNode.children = fullNode.children[b.minDeg:]
	}
	

	fullNode.nk = b.minDeg - 1
	n.insertChild(i+1,nNode)
	n.insertKey(i, fullNode.keys[b.minDeg - 1])
	fullNode.keys = fullNode.keys[:b.minDeg-1]
	fullNode.children = fullNode.children[:b.minDeg]
	//write? fullNode, nnode and n
}

func (b *BTree) mergeChild(n *Node, i int) {
	//What is we are at last child
	child, sibling := n.children[i], n.children[i+1]

	for i, key := range sibling.keys {
		child.keys[(b.minDeg-1)+i] = key
	}

	if !child.isLeaf {
		for i, child := range sibling.children {
			child.children[b.minDeg+i] = child
		}
	}

}

func (n *Node) redistributeChild(child *Node) bool{
	var childpos int
	var lsib, rsib *Node
	
	for _, key := range n.keys {
		if child.keys[0] > key {
			childpos++
		}
	}

	if childpos > 0 {
		lsib = n.children[childpos-1]
	}

	if childpos < len(n.children)-1 {
		rsib = n.children[childpos+1]
	}

	//take sibling key make parent separator / place parent separater within child
	// Adjust children?
	// Q How do you insert keys 
	if lsib != nil && len(lsib.keys) > n.mindeg-1{
		l := len(lsib.keys)-1
		pk := []string{n.keys[childpos-1]}
		lk := lsib.keys[l]
		
		lsib.keys = lsib.keys[:l]
		n.keys[childpos-1] = lk
		//append to beginning
		child.keys = append(pk,child.keys...) 

		//siblings child given to child
		if !lsib.isLeaf {
			n.children[childpos-1] = lsib.children[l]
		}
		return true
	} else if rsib != nil && len(rsib.keys) > n.mindeg-1 {
		rk := rsib.keys[0]
		pk := n.keys[childpos+1]

		rsib.keys = rsib.keys[1:]
		n.keys[childpos+1] = rk
		child.keys = append(child.keys, pk)

		if !rsib.isLeaf {
			n.children[childpos+1] = rsib.children[0]
		}
		return true
	}

	return false
}

func (b *BTree) findMerge(key string, pn *Node) {
	var cPos int
	for k,v := range pn.children {
		if v.keys[0] == key {
			cPos = k
		}
	}
	b.mergeChild(pn, cPos)
}
// TODO: check the correctness of this function
func display(node *Node) {
	if node == nil {
		return
	}

	var i int
	for i = 0; i < node.nk; i++ {
		if !node.isLeaf {
			display(node.children[i])
		}
		fmt.Println(node.keys[i])
	}
}
