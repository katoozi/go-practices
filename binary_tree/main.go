package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree is a struct that we used for implementing binary tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// traverse will print all tree nodes
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

// create a binary tree with custom number of nodes
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

// insert new node to the tree.
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}

	// balancing operations
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("root value: ", tree.Value)
	traverse(tree)
	fmt.Println()
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("root value: ", tree.Value)
}
