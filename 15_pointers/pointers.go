package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("Change num", num)
// }

func changeNum(num *int) {
	*num = 5
	fmt.Println("Change num", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)

	// fmt.Println("Memory Address", &num)

	fmt.Println("After change num", num)
}
