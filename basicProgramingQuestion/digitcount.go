package main
import "fmt"
func main(){
	var input, count int
	fmt.Println("Enter a number:")
	fmt.Scan(&input)
	res := input
	count = 0
	for input != 0 {
		count+=1
		input = input / 10
	}
	fmt.Printf("Number of digits in %d is: %d\n",res,count)
}