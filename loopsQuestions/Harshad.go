package main
import "fmt"
func main(){
	var num int
	fmt.Print("Enter a Number: ")
	fmt.Scan(&num)
	sum := 0
	temp := num
	for num !=0 {
		rem := num % 10
		sum += rem
		num /= 10
	}
	if temp % sum == 0{
		fmt.Printf("%d is a Harshad Number",temp)
	}else{
		fmt.Printf("%d is not a Harshad Number",temp)
	}
}