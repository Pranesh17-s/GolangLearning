package main
import "fmt"
func main(){
	var Range int
	fmt.Print("Enter a Range: ")
	fmt.Scan(&Range)
	if Range < 1{
		fmt.Println("Invalid Range, Range should be greater than 0")
	}
	oddSum := 0
	evenSum := 0
	for i:=1; i<=Range; i++{
		if i%2==0{
			evenSum += i
		}else{
			oddSum += i
		}
	}
	fmt.Printf("Odd Sum from 1 to %d is %d\n",Range,oddSum)
	fmt.Printf("Even Sum from 1 to %d is %d",Range,evenSum)
}