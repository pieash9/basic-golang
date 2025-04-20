package main

import "fmt"

// iterating over data structure
func main() {
	// nums := []int{6,7,8}

	// sum:=0

	// for i, num := range nums {
	// 	fmt.Println(i, num)
	// }

	// m:= map[string]string{"fname":"pieash", "lname":"khan"}

	// for k,v := range m {
	// 	fmt.Println(k,v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// unicode point rune
	// starting byte rune
	// 300 -> 1 byte , 2 byte
	for i,c := range "পিয়াশ" {
		fmt.Println(i,string(c))
	}

}