package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, dict
func main() {
	// creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exist in the map then it return zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50

	// fmt.Println(m["phone"])

	// delete(m, "price")
	// clear(m)

	// fmt.Println((m))

	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}

	// v, ok := m["phones"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 2}

	fmt.Println(maps.Equal(m1, m2))
}
