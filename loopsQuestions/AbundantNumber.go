package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	sum := 1
	for i := 2; i <= input/2; i++ {
		if input%i == 0 {
			sum += i
		}
	}
	if sum > input {
		fmt.Printf("%d is an Abundant Number\n", input)
	} else {
		fmt.Printf("%d is not an Abundant Number\n", input)
	}
}
