package main

import "fmt"

func main(){

	var confName ="Learning Go"
	const noOfTickets = 64
	var remTickets = 64
	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Printf("We have %v tickets and %v tickets remaining, Hurry UP!!\n",noOfTickets,remTickets)
	fmt.Println("Get Your tickets")
}