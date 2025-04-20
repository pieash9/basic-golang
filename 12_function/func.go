package main

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguage ()(string, string,string){
// 	return "golang" , "javascript", "c"
// }

// func processIt ( fn func(a int)int) {
// 	fn(1)
// }

func processIt() func(a int)int {
	return func (a int)int{
		return 4
	}
}

func main() {
	// fn:= func(a int) int {
	// 	return 2
	// }

	fn := processIt()
	fn(6)

	// processIt(fn)
	// lang1, lang2, _ := getLanguage()
	// fmt.Println(lang1,lang2)
	// result := add(4, 5)
	// fmt.Println(result)
}