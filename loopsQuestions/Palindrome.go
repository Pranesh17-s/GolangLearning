package main
import "fmt"
func main(){
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	var res int
	temp := input
	for input != 0{
		res = res*10 + input % 10
		input /= 10
	}
	if res == temp{
		fmt.Printf("%d is a Palindrome number\n",temp)
	}else{
		fmt.Printf("%d is not a Palindrome number\n",temp)
	}
}