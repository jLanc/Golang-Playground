package binarysearchtree

import (
	"reflect"
	"testing"
)

var bst Tree

func fillTree(bst *Tree) {
	bst.Insert(8, 8)
    bst.Insert(4, 4)
    bst.Insert(10, 10)
    bst.Insert(2, 2)
    bst.Insert(6, 6)
    bst.Insert(1, 1)
    bst.Insert(3, 3)
    bst.Insert(5, 5)
    bst.Insert(7, 7)
    bst.Insert(9, 9)
}

func isSameSlice(a Node, b Node) bool {
   return reflect.DeepEqual(a,b)
}

func TestInsert(t *testing.T) {
	fillTree(&bst)
	bst.PrintTree()

	bst.Insert(11, 11)
	bst.PrintTree()
}

func TestSearch(t *testing.T) {
    // Setup the tree
    fillTree(&bst)

    if !bst.Search(1) || !bst.Search(8) || !bst.Search(5) {
        bst.PrintTree()
        t.Error("Failed to find elements 1, 8 and 5 which are know to be ins tree")
    }
}