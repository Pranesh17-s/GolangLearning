package main
import "fmt"
func main() {
	// Integer Types in Go
	// int8	8-bit signed integer
	// int16	16-bit signed integer
	// int32	32-bit signed integer
	// int64	64-bit signed integer
	// uint8	8-bit unsigned integer
	// uint16	16-bit unsigned integer
	// uint32	32-bit unsigned integer
	// uint64	64-bit unsigned integer
	// int	Both int and uint contain same size, either 32 or 64 bit.
	// uint	Both int and uint contain same size, either 32 or 64 bit.
	// rune	It is a synonym of int32 and also represent Unicode code points.
	// byte	It is a synonym of uint8.
	// uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
	var intVar int = 42
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807
	fmt.Println("Integer Types:")
	fmt.Println("int:", intVar)
	fmt.Println("int8:", int8Var)
	fmt.Println("int16:", int16Var)
	fmt.Println("int32:", int32Var)
	fmt.Println("int64:", int64Var)
	fmt.Println()
	// Basic Arithmetic Operations
	var x int16 = 170
    var y int16 = 83
    //Addition
    fmt.Printf(" addition :  %d + %d = %d\n ", x, y, x+y)
    //Subtraction
    fmt.Printf("subtraction : %d - %d = %d\n", x, y, x-y)
    //Multiplication
    fmt.Printf(" multiplication : %d * %d = %d\n", x, y, x*y)
    //Division
    fmt.Printf(" division : %d / %d = %d\n", x, y, x/y)
    //Modulus
    fmt.Printf(" remainder : %d %% %d = %d\n", x, y, x%y)
	fmt.Println()

	// Floating-Point Types in Go
	// float32	32-bit IEEE 754 floating-point number
	// float64	64-bit IEEE 754 floating-point number

	a := 20.45
    b := 34.89
    
    // Subtraction of two 
    // floating-point number
    c := b-a
    
    // Display the result 
    fmt.Printf("Result is: %f", c)
    
    // Display the type of c variable
    fmt.Printf("\nThe type of c is : %T", c)

	fmt.Println()

	// Basic float arithmetic operations

	var n float32 = 5.00
    var m float32 = 2.25
    //Addition
    fmt.Printf("addition :  %g + %g = %g\n ", n, m, n+m)
    //Subtraction
    fmt.Printf("subtraction : %g - %g = %g\n", n, m, n-m)
    //Multiplication
    fmt.Printf("multiplication : %g * %g = %g\n", n, m, n*m)
    //Division
    fmt.Printf("division : %g / %g = %g\n", n, m, n/m)
	fmt.Println()

	// Boolean Type in Go
	// The boolean type in Go is represented by the keyword bool. It can hold one of two values: true or false.
	var boolVar1 bool = true
	var boolVar2 bool = false
	fmt.Println("Boolean Types:")
	fmt.Println("boolVar1:", boolVar1)
	fmt.Println("boolVar2:", boolVar2)
	fmt.Println()

	// String Type in Go
	// The string type in Go is represented by the keyword string. Strings are sequences of characters and are immutable.
	var strVar string = "Hello, Go!"
	fmt.Println("String Type:")
	fmt.Println("strVar:", strVar)
	fmt.Println("Length of strVar:", len(strVar))
	fmt.Println("Character at index 1:", string(strVar[1]))
	fmt.Println()

	// Formatted Output
	fmt.Print("Formatted Output:\n")
	fmt.Printf("Integer: %d\n", intVar)
	fmt.Printf("Float: %f\n", a)
	fmt.Printf("Boolean: %t\n", boolVar1)
	fmt.Printf("String: %s\n", strVar)
	fmt.Printf("Type of intVar: %T\n", intVar)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of boolVar1: %T\n", boolVar1)
	fmt.Printf("Type of strVar: %T\n", strVar)
}