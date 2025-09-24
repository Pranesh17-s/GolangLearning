package main
import "fmt"
import "math"
func main(){
	var num int
	fmt.Print("Enter a Number: ")
	fmt.Scan(&num)
	digits := int(math.Log10(float64(num)))+1
	sum := 0
	temp := num
	for num != 0{
		rem := num % 10
		sum += int(math.Pow(float64(rem),float64(digits)))
		num /= 10
	}
	if sum == temp{
		fmt.Printf("%d is an Armstrong Number\n",temp)
	}else{
		fmt.Printf("%d is not an Armstrong Number\n",temp)
	}
}