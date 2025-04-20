package main

func main() {
	// simple switch

	// i := 1

	// switch i {
	// case 1:
	// 	println("1")
	// case 2:
	// 	println("2")
	// case 3:
	// 	println("3")
	// default:
	// 	println("default")
	// }

	// multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("weekend")
	// default:
	// 	println("Work day")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("int")
		case string:
			println("string")
		case bool:
			println("bool")
		default:
			println("unknown")
		}
	}

	whoAmI(5)

}
