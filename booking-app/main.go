package main

import "fmt"

func main(){

	var confName ="Learning Go"
	const noOfTickets = 64
	var remTickets uint = 64
	// fmt.Printf("Type of confName is %T,noOfTickets is %T,remTickets is %T\n",confName,noOfTickets,remTickets)

	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Printf("We have %v tickets and %v tickets remaining, Hurry UP!!\n",noOfTickets,remTickets)
	fmt.Println("Get Your tickets")
	
	// var bookings = [64]string{}
	// booking := []string{}

	for{
			var booking []string
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no of Tickets: ")
	fmt.Scan(&userTickets)

	// Pointer
	// fmt.Println(userName)
	// fmt.Println(&userName)

	// booking[0] = firstName + " " + lastName	

	booking = append(booking,firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remTickets,confName)

	fmt.Println("The bookings are: ",booking)
	}
}