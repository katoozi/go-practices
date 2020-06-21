package main

import "fmt"

func main() {
	firstMap := make(map[int]string)
	secondMap := map[int]string{
		1: "Jone Doe",
		2: "mohammad katoozi",
	}
	fmt.Println(firstMap)
	fmt.Println(secondMap)
	fmt.Println("second map 1", secondMap[1])
	fmt.Println("second map 2", secondMap[2])

	fmt.Println("Start Iterating")
	for key, value := range secondMap {
		fmt.Println("key: ", key, " value: ", value)
	}
	fmt.Println("End Iterating")

	delete(secondMap, 1)
	delete(secondMap, 1)
	delete(secondMap, 1)
	delete(secondMap, 1) // delete will not return error is the given key doesn't exists.
	fmt.Println("after delete", secondMap)

	_, ok := secondMap[2]
	if ok {
		fmt.Println("Key 2 is exists")
	} else {
		fmt.Println("Key 2 is not exists")
	}
}
