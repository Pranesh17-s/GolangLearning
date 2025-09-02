package main

import "fmt"

func main(){

	var confName ="Learning Go"
	const noOfTickets = 64
	var remTickets = 64
	// fmt.Printf("Type of confName is %T,noOfTickets is %T,remTickets is %T\n",confName,noOfTickets,remTickets)
	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Printf("We have %v tickets and %v tickets remaining, Hurry UP!!\n",noOfTickets,remTickets)
	fmt.Println("Get Your tickets")
	var userName string
	var userTickets int
	userName = "Pranesh"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n",userName,userTickets)
}