package main

import "fmt"

func main() {
	// This is a introduction on loops in go
	// Go has only one loop that is for loop you can use that as while loop.
	// for i:=0;i<5;i++{
	// 	fmt.Printf("%v\n",i)
	// }
	// fmt.Println("Loop ended")

	//For infinite loop
	// for{
	// 	fmt.Println("This is infinite loop")
	// }

	// Introduction to use for loop as while loop
	i := 0
	for i < 6 {
		fmt.Printf("%d\n", i)
		i++
	}
	fmt.Println("Loop ended")
}
