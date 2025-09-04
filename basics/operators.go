package main
import "fmt"
func main(){
	// Arithmetic Operators
	a := 52
	b := 64
	fmt.Printf("Addition of %d and %d is %d\n",a,b,a+b)
	fmt.Printf("Subtraction of %d and %d is %d\n",a,b,a-b)
	fmt.Printf("Multiplication of %d and %d is %d\n",a,b,a*b)
	fmt.Printf("Division of %d and %d is %d\n",a,b,a/b)
	fmt.Printf("Modulus of %d and %d is %d\n",a,b,a%b)
	fmt.Println()

	// Relational Operators

	j := 52
	s := 64

	fmt.Println(j==s)
	fmt.Println(j!=s)
	fmt.Println(j<s)
	fmt.Println(j>s)
	fmt.Println(j<=s)
	fmt.Println(j>=s)
	
	// Logical Operators
	firstNum := 10
	secondNum := 20
	if(firstNum != secondNum && firstNum<=secondNum){
		fmt.Println("True")
	}
	if(firstNum == secondNum || secondNum > firstNum){
		fmt.Println("True")
	}
	if(!(firstNum == secondNum)){
		fmt.Println("True")
	}
}