package binarysearchtree

import (
	"fmt"
	"sync"
)

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *Tree) Insert(key int, value int) {
	//Lock tree as this is critical section
	bst.lock.Lock()
	defer bst.lock.Unlock()

	//create new node
	n := &Node{key, value, nil, nil}

	//set root if not set
    if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// recursively traverse tree to insert into correct position
func insertNode(node, newNode *Node) {
	// check insert left
	if newNode.key < node.key {
		// only insert into left if it's empty
		if node.left == nil {
			node.left = newNode
		} else {
			//recursively call this function to continue traversal insert
			insertNode(node.left, newNode)
		}		
	} else {
		//assume value should be on the right
		// only insert into right if it's empty
		if node.right == nil {
			node.right = newNode
		} else {
			// recursively call this function to continue traversal insert
			insertNode(node.right, newNode)
		}
	}
}

// Search returns a true if item is found
func (bst *Tree) Search(key int) bool {
	// read lock the tree
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return search(bst.root, key)
}

// internal function to  recursively traverse tree to search if number exists
func search(n *Node, key int) bool {
	// test if we've reached the final node
	if n == nil {return false}

	// contrinue traversing 
	if key < n.key {
		search (n.left, key)
	}
	if key > n.key {
		search (n.right, key)
	}
	return true
}

/*func depthFirstSearch(n *Node, key int) bool {

}*/

func (bst *Tree) PrintTree() {
	// Lock tree so it can't change when printing / traversing 
	bst.lock.Lock()
	defer bst.lock.Unlock()

	// Print tree
	fmt.Println("------------------------------------------------")
    printTreeUtil(bst.root, 0)
    fmt.Println("------------------------------------------------")
}

func printTreeUtil (n *Node, level int) {
	// Start by checking if node exists
	if n != nil {
		// Add spacing to the node depth given the level 
		format := ""
		for i := 0; i < level; i++ {
			format += "     "
		}
		
		format += "---["
		// iterate on level
		level ++

		// call recursively to traverse tree
		printTreeUtil(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		printTreeUtil(n.right, level)
	}
}