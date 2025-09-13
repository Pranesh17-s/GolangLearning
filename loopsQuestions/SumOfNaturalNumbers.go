package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	if input < 1 {
		fmt.Println("Invalid Input")
		return
	}
	res := 0
	for i := 1; i < input; i++ {
		res += i
	}
	fmt.Printf("Sum of first %d natural numbers is %d\n", input, res)
}
