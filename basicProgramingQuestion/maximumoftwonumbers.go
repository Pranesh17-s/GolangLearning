package main
import "fmt"
func main(){
	var number1 int
	fmt.Println("Enter first number:")
	fmt.Scan(&number1)
	var number2 int
	fmt.Println("Enter second number:")
	fmt.Scan(&number2)
	if(number1>number2){
		fmt.Printf("%d is maximum\n",number1)
	}else if(number1<number2){
		fmt.Printf("%d is maximum\n",number2)
	}else{
		fmt.Print("Both are equal")
	}
}