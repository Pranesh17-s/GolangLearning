package main

import "fmt"

func main() {
	var number1 int
	fmt.Println("Enter first number:")
	fmt.Scan(&number1)
	var number2 int
	fmt.Println("Enter second number:")
	fmt.Scan(&number2)
	var number3 int
	fmt.Println("Enter third number:")
	fmt.Scan(&number3)

	if number1 >= number2 && number1 >= number3 {
		fmt.Printf("%d is maximum\n", number1)
	} else if number2 >= number1 && number2 >= number3 {
		fmt.Printf("%d is maximum\n", number2)
	} else if number3 >= number1 && number3 >= number2 {
		fmt.Printf("%d is maximum\n", number3)
	} else {
		fmt.Print("All are equal")
	}
}
