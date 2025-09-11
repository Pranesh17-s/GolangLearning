package main
import "fmt"
func main(){
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	res := 0
	temp := input
	for input != 0{
		res += input % 10
		input /= 10
	}
	fmt.Printf("Sum of digits of number %d is %d\n",temp, res)
}