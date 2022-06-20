package main

// package declaration

import (
	"fmt"
	"math/rand"
	"time"
)

// import function

func main() {
	fmt.Println("guessing game started...")
	var userInput int
	var secretNumber int
	// var quit bool = false
	quit := false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)

	for quit != true {

		// take user input
		fmt.Printf("Please enter a number: ")
		// save the user input using scan
		fmt.Scan(&userInput)
		// print the number
		fmt.Println("Your given number is:", userInput)

		if userInput == secretNumber {
			fmt.Println("You won!")
			quit = true
		} else if userInput < secretNumber {
			fmt.Println("Too low!!")
		} else if userInput > secretNumber {
			fmt.Println("Too High!!!")
		}

	}

}
