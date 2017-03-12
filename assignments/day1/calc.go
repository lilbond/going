package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a - b
}

func divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println("Sum of 10, 2 is : " + strconv.Itoa(add(10, 2)))
	fmt.Println("Difference of 10, 2 is : " + strconv.Itoa(subtract(10, 2)))
	fmt.Println("Multiplication of 10, 2 is : " + strconv.Itoa(mul(10, 2)))
	fmt.Println("Division of 10, 2 is : " + strconv.Itoa(divide(10, 2)))
}
