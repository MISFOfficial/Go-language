package main

import "fmt"

// (int, int) as defing the type means i want to return two value integer type

func addition(num1 int, num2 int) (int int) {
	return num1 + num2
}

func main() {
	fmt.Println("Funtions and Return Value and Type")
	sum := addition(10, 20)
	fmt.Println("Addition result:", sum)

	temp, rainfall := CelsiusToFahrenheit(25)
	fmt.Println("Temperature:", temp, "Fahrenheit")
	fmt.Println("Rainfall:", rainfall, "Inches")
}

func CelsiusToFahrenheit(celsius float64) (float64, float64) {
	fahrenheit := (celsius * 9 / 5) + 32
	rainfall := celsius * 0.0393701 // Convert mm to inches
	return fahrenheit, rainfall
}
