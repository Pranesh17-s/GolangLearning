package main
import "fmt"
func main() {
	var a int = 10
	if(a>9){
		fmt.Printf("%d is greater than 9\n",a)
	}
	fmt.Printf("%d is a two digit number\n",a)
	// If-else statement
	if(a>9){
		fmt.Printf("%d is greater than 9\n",a)
	}else{
		fmt.Printf("%d is not greater than 9\n",a)
	}
	// Nested if statement
	if(a>9){
		if(a<100){
			fmt.Printf("%d is a two digit number\n",a)
		}
	}
	// if-else-if ladder
	if(a<0){
		fmt.Printf("%d is a negative number\n",a)
	}else if(a==0){
		fmt.Printf("%d is zero\n",a)
	}else{
		fmt.Printf("%d is a positive number\n",a)
	}
	//if else-if ladder to find the largest of three numbers
	var b int = 20
	var c int = 15
	if(a>=b && a>=c){
		fmt.Printf("%d is the largest number\n",a)
	}else if(b>=a && b>=c){
		fmt.Printf("%d is the largest number\n",b)
	}else if(c>=a && c>=b){
		fmt.Printf("%d is the largest number\n",c)
	}else{
		fmt.Printf("All numbers are equal\n")
	}
}