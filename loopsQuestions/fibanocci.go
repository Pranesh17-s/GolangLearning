package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	first := 0
	second := 1
	third := 0
	fmt.Printf("Fibonacci Series: %d %d ", first, second)
	for i := 2; i < input; i++ {
		third = first + second
		fmt.Printf("%d ", third)
		first = second
		second = third
	}
}
