package main

import "fmt"

const name = "pieash1"
const age = 31
func main() {
	// const name = "pieash"
	// const age = 30
	// println(name, age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}