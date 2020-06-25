package main

import "fmt"

// Node is a element of linked list
type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

var root = new(Node)

// addNode will add new node to the end of linked list
func addNode(t *Node, v int) int {
	if root == nil {
		root = &Node{v, nil, nil}
		return 0
	}
	if v == t.Value {
		fmt.Println("Node Already in list.")
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil, t}
		return -2
	}
	return addNode(t.Next, v)
}

// traverse will print all linked list elements
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List")
		return
	}

	for t != nil {
		fmt.Printf("-> %d", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// reverse will print all linked list elements in reverse direction
func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List")
		return
	}
	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("-> %d", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("-> %d", temp.Value)
	fmt.Println()
}

// lookupNode will search for an node in linked list
func lookupNode(t *Node, v int) bool {
	if t == nil {
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

// size will calculate linked list size
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty List")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i

}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	reverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	reverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 55)
	addNode(root, 555)
	addNode(root, 444)
	traverse(root)
	reverse(root)
	addNode(root, 100)
	traverse(root)
	reverse(root)

	if lookupNode(root, 100) {
		fmt.Println("100 is exists in linked list")
	} else {
		fmt.Println("100 is not exists in linked list")
	}
	if lookupNode(root, 101) {
		fmt.Println("101 is exists in linked list")
	} else {
		fmt.Println("101 is not exists in linked list")
	}
	fmt.Println("Size of linked list: ", size(root))
}
