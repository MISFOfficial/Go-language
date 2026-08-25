package main

import "fmt"

func main() {
	a := 10
	// fmt.Println(a)

	if a > 100 {
		fmt.Println(false)
	} else if a == 100 {
		fmt.Println(true)
	} else {
		fmt.Println("Tmr dara programming hobe na")
	}

	switch a {
	case 100:
		fmt.Println("This is 100")
	case 10:
		fmt.Println("This is 10")
	default:
		fmt.Println("This is default")
	}
}
