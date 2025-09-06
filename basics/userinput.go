package main
import "fmt"
func main(){
	// Get user input for int
	var number int
	fmt.Scan(&number)
	fmt.Print(number)

	// Get user input for float
	var floatNum float64
	fmt.Scan(&floatNum)
	fmt.Print(floatNum)
	
	// Get user input for string
	var str string
	fmt.Scan(&str)
	fmt.Print(str)

	// Get user input for boolean
	var boolean bool
	fmt.Scan(&boolean)
	fmt.Print(boolean)

	// Get user input for multiple values
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(a, b)

	// Get user input for multiple values of different types
	var name string
	var age int
	fmt.Scan(&name, &age)
	fmt.Print(name, age)

	// Get user input using fmt.Scanf
	var c int
	fmt.Scanf("%d", &c)
	fmt.Print(c)

	// Get user input using fmt.Scanln
	var d int
	fmt.Scanln(&d)
	fmt.Print(d)

	// Get user input using bufio package
	// Uncomment the following lines to use bufio for user input
	/*
	import (
		"bufio"
		"os"
	)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Print(input)
	*/

}