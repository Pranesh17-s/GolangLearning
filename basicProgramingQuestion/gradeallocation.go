package main
import "fmt"
func main(){
	var mark int
	fmt.Scan(&mark)
	if(mark>=90 && mark<=100){
		fmt.Print("A+")
	}else if(mark>=80 && mark<90){
		fmt.Print("A")
	}else if(mark>=70 && mark<80){
		fmt.Print("B+")
	}else if(mark>=60 && mark<70){
		fmt.Print("B")
	}else if(mark>=50 && mark<60){
		fmt.Print("C")
	}else{
		fmt.Print("Sorry you didn't pass the exam")
	}
}