package main

import "fmt"

// SIZE is the cardinalities of hash table
const SIZE = 15

// Node is a element in hash table
type Node struct {
	Value int
	Next  *Node
}

// HashTable is the hash table structure
type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

// insert into hash table
func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	node := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &node
	return index
}

// print hash table
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

// detemine that given element already exists in hash table
func lookup(hash *HashTable, value int) bool {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				if t.Value == value {
					return true
				}
				t = t.Next
			}
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces: ", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
	fmt.Println("exists: ", 10, lookup(hash, 10))
	fmt.Println("exists: ", 1000, lookup(hash, 1000))
}
