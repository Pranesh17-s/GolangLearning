package main
import "fmt"
func main(){
	var mark1,mark2,mark3,mark4,mark5 int
	fmt.Println("Enter marks of five subjects:")
	fmt.Scan(&mark1,&mark2,&mark3,&mark4,&mark5)
	totalmarks:=mark1+mark2+mark3+mark4+mark5
	percentage:=(totalmarks*100)/500
	fmt.Printf("Total marks is %d\n",totalmarks)
	fmt.Printf("Percentage is %d\n",percentage)
	if percentage >= 90 {
		fmt.Print("Grade is A+")
	} else if percentage >= 80 {
		fmt.Print("Grade is A")
	} else if percentage >= 70 {
		fmt.Print("Grade is B+")
	} else if percentage >= 60 {
		fmt.Print("Grade is B")
	} else if percentage >= 50 {
		fmt.Print("Grade is C")
	} else {
		fmt.Print("Sorry you didn't pass the exam")
	}