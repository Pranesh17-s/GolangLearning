package main
import(
	"fmt"
	"math"
)
func main(){
	var a,b,c float64
	fmt.Println("Enter coefficients a, b and c:")
	fmt.Scan(&a,&b,&c)
	discriminant:=b*b-4*a*c
	if(discriminant>0){
		root1:=(-b+math.Sqrt(discriminant))/(2*a)
		root2:=(-b-math.Sqrt(discriminant))/(2*a)
		fmt.Printf("Roots are real and different.\nRoot 1: %.2f\nRoot 2: %.2f\n",root1,root2)
	}else if(discriminant==0){
		root:=(-b)/(2*a)
		fmt.Printf("Roots are real and same.\nRoot: %.2f\n",root)
	}else{
		realPart:=-b/(2*a)
		imaginaryPart:=math.Sqrt(-discriminant)/(2*a)
		fmt.Printf("Roots are complex and different.\nRoot 1: %.2f + %.2fi\nRoot 2: %.2f - %.2fi\n",realPart,imaginaryPart,realPart,imaginaryPart)
	}
}