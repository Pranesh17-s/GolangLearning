package main
import "fmt"
func main(){
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	if input <= 1{
		fmt.Printf("%d is not a Prime Number\n",input)
		return
	}
	for i:=2; i*i<=input;i++{
		if(input % i == 0){
			fmt.Printf("%d is not a Prime Number\n",input)
			return
		}
	}
	fmt.Printf("%d is a Prime Number\n",input)
}