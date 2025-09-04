package main  //-> package declaration because every go file is part of a package
import "fmt" //-> import declaration to use other packages and fmt is a package to format text, including printing to the console.

func main(){  //-> main function is the entry point of a Go program. Execution starts from here.
	fmt.Println("Hello World")  //-> Print function from fmt package is used to print text to the console this added a new line at the end.
	fmt.Print("New day new Learning") //-> Print function from fmt package is used to print text to the console without a new line at end.
}