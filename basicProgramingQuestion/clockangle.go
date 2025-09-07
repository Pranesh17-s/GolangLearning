package main
import "fmt"
func main(){
	var hour,minute int
	fmt.Println("Enter hour and minute:")
	fmt.Scan(&hour,&minute)
	hourAngle:=float64(hour%12)*30+float64(minute)*0.5
	minuteAngle:=float64(minute)*6
	angle:=hourAngle-minuteAngle
	if(angle<0){
		angle=-angle
	}
	if(angle>180){
		angle=360-angle
	}
	fmt.Printf("Angle between hour and minute hand: %.2f degrees\n", angle)
}