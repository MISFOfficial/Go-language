package main

import "fmt"

func addition(num1 int, num2 int) (int int) {
	return num1 + num2
}

func main() {
	fmt.Println("Funtions and Return Value and Type")
	sum := addition(10, 20)
	fmt.Println("Addition result:", sum)
}
