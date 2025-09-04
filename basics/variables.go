package main
import "fmt"

func main(){
	var name = "Pranesh" //-> 'var' keyword is used to declare a variable, 'name' is the variable name, and it is assigned the string value "Pranesh".
	fmt.Println(name)
	const value = 10
	fmt.Println(value)
	// value = 20  //-> This will cause an error because value is a constant and cannot be changed.
	age := 21 //-> Short variable declaration using ':=' operator, which infers the type from the assigned value.
	fmt.Println(age)
}