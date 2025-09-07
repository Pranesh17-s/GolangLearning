package main
import "fmt"
func main(){
	var principal,rate,time float64
	fmt.Println("Enter principal amount, rate of interest and time in years:")
	fmt.Scan(&principal,&rate,&time)
	simpleInterest:=(principal*rate*time)/100
	fmt.Printf("Simple Interest: %.2f\n",simpleInterest)
}