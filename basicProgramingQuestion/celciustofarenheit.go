package main
import "fmt"

func main() {
    var celsius float64
    fmt.Print("Enter temperature in Celsius: ")
    fmt.Scan(&celsius)
    fahrenheit := (celsius * 9/5) + 32
    fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheit)

	var fahrenheitInput float64
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scan(&fahrenheitInput)
	celsiusConverted := (fahrenheitInput - 32) * 5/9
	fmt.Printf("Temperature in Celsius: %.2f\n", celsiusConverted)
}
