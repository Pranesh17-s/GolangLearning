package main
import "fmt"
import "math"
func main(){
	var num int
	fmt.Print("Enter a Number: ")
	fmt.Scan(&num)
	temp := num
	square := num * num
	digit := int(math.Log10(float64(num))) + 1
	mod := int(math.Pow(10,float64(digit)))
	modSquare := square % mod
	if modSquare == temp{
		fmt.Printf("%d is a Automorphic Number",temp)
	}else{
		fmt.Printf("%d is not a Automorphic Number",temp)
	}
}