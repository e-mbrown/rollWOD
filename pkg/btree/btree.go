package btree

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
		children: []*Node,
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

// TODO: Write test
func (b *BTree) search(n *node, key string) (*Node, int){
	low, high := 0, n.nk

	for low < high {
		mid = low + (high-low)/2
		if key == n.keys[mid]{
			return n, i
		}else if key < n.keys[mid] {
			high = mid
		}else {
			low = mid+1
		}
	}

	if low < n.nk && key == n.key[low] {
		return n,i
	}

	if n.isLeaf{
		//key doesnt exist so maybe return error type
		return nil, nil
	}

	return b.search(n.children[low], key)
}

// I think we can just return the node?
func (b *BTree) searchInsertPos(n *Node, key int) (*Node, int){
	low, high := 0, n.nk

	for low < high {
		mid = low + (high-low)/2
		if key == n.keys[mid]{
			// Duplicate key, error?
			return nil,nil
		}else if key < n.keys[mid] {
			high = mid
		}else {
			low = mid+1
		}
	}

	if low < n.nk && key == n.key[low] {
		// Duplicate
		return nil, nil
	}

	if n.isLeaf{
		// return and insert in out func
		return n, i
	}	
	
	if n.children[low].nk == 2*b.minDeg-1 {
		b.splitChild(n, low)
		if key > n.keys[low] { low++ }
	}

	return b.searchInsertPos(n.children[low], key)
}
	/*
 Insert
 search for leaf node and insert
 If node full
 	Split node into two at median
	keys to the left of median go to 1st
	keys to the right go to 2nd
	Median moves up to parent
 If parent full
 	split parent up to root
*/
func (b *BTree) insert(key string){
	root := b.root
	if b.IsFull(root) {
		root = b.SplitRoot()
	}
	
	b.insertNotFull(root, key)
}

func (b *BTree) insertNotFull(n *Node, k string){
	//theoretically node is a leaf and i is
	// key that our key may or may not be larger than
	node, i := n.searchInsertPos(n.root, k)
	var left []*Node
	if key < n.keys[i]{
		copy(left,n.keys[:i])
		left = append(left, key)
		left = append(left, n.keys[i:])
	} else {
		copy(left, n.keys[:i+1])
		left = append(left, key)
		left = append(left, n.keys[i+1:])
	}
	// disk write?
}

//case: reach leaf node, reach internal contains target,
// reach internal no target
// if root is deleted? and has children 1 vs more
func(b *BTree) delete(n *Node, key string) {	
	tNode, kpos := b.search(n, key)
	//if err no key

	if tNode.isLeaf {
		n.keys = append(n.keys[:i], n.keys[i+1]...)
		return
	}

	if len(tNode.children[kpos].keys) >= self.t {
		predk := b.getPred(tNode, kpos)
		b.delete(tNode.children[kpos], predk)
		tNode.keys[kpos] = predk
	} else if len(tNode.children[kpos+1].keys) >= self.t {
		succk := b.getSucc(tNode, kpos)
		b.delete(tNode.children[kpos+1],succk)
		tNode.keys[kpos] = succk
	} else {

		
func (b *Btree) getPred(n *Node, idx int) string {
	curr := n.children[idx]

	for !curr.isLeaf {
		curr = curr.children[curr.nk]
	}

	return curr.keys[curr.nk-1]
}


func (b *Btree) getSucc(n *Node, idx int) string {
	curr := n.children[idx+1]

	for !curr.isLeaf {
		curr = curr.children[0]
	}

	return curr.keys[0]

func (b *BTree) splitRoot() {
	nRoot = b.createNode(false)
	nRoot.children = append(nRoot.children, b.root)
	b.root = nRoot
	b.splitChild(nRoot, 0)
	return
}

func (b *BTree) splitChild(n *Node, i int) {
	fullNode = n.children[i]
	nNode := createNode(fullNode.isLeaf)
	nNode.nk = b.minDeg - 1
	//Taking greater keys
	nNode.keys = fullNode.keys[b.minDeg:]
	
	if !fullNode.isLeaf{
		nNode.children = fullNode.children[b.minDeg:]
	}
	

	fullNode.nk = b.minDeg - 1
	n.insertChild(i+1,nnode)
	n.insertKey(i, fullNode.keys[b.minDeg - 1])
	fullNode.keys = fullNode.keys[:b.minDeg-1]
	fullNode.children = fullNode.children[:b.minDeg]
	//write? fullNode, nnode and n
}

func (b *BTree) mergeChild(n *Node, idx int) {
	child, sibling := node.children[i], node.children[i+1]

	child.keys[b.minDeg-1] = sibling.keys[idx]
	for i, key := range sibling.children {
		child.keys[b.minDeg+i] = key
	}

	if !child.isLeaf {
		for i := range sibling.nk + 1{
			child.children[b.minDeg+i] = sibling.children[i]
		}
	}


// TODO: check the correctness of this function
func display(node *Node) {
	if node == NULL {
		return
	}

	var i int
	for i = 0; i < node.nc; i++ {
		if !node.leaf {
			display(node.children[i])
		}
		fmt.Println(node.keys[i])
	}
}
