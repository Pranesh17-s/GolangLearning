package main 
import "fmt"
func main(){
	var start, end int
	fmt.Print("Enter the range (start end): ")
	fmt.Scan(&start, &end)
	if start > end || start < 0 || end < 0{
		fmt.Println("Invalid Range")
		return
	}
	if end < 2{
		fmt.Printf("No Prime Numbers in the range %d to %d\n",start,end)
		return
	}
	for num := start; num <= end; num++{
		isPrime := true
		for i:=2; i*i<=num;i++{
			if(num % i == 0){
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(num, " ")
		}
	}
}