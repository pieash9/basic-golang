package main

import "fmt"

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// type stack[T any] struct {
// 	elements []T
// }

func main() {
	// myStack := stack[string]{
	// 	elements: []string{"test", "2", "3"},
	// }

	// fmt.Println(myStack.elements)
	// nums := []int{1, 2, 3, 4}
	// names := []string{"pieash", "khan", "khan2"}
	vals := []bool{true, false, true}
	// printStringSlice(names)
	printSlice(vals, "pieash")
}
