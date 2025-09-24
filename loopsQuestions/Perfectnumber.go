package main
import "fmt"
import "math"
func main(){
	var num int
	fmt.Print("Enter a Number: ")
	fmt.Scan(&num)
	b := math.Sqrt(float64(num))
	val := int(b*b)
	if(val == num){
		fmt.Printf("%d is a Perfect Number\n",num)
	}else{
		fmt.Printf("%d is not a Perfect Number\n",num)
	}
}