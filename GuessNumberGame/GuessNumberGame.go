package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var randomNum = rand.Intn(100)
	fmt.Println("Let's play a number guessing game! Guess a number between 0-100 and I'll say if it's higher or lower:") //Start game
	var attemptsValue int = checkNum(randomNum)                                                                          //Call function, return amount of attempts it took
	fmt.Printf("Congrats! The number was indeed %v! It only took %v attempts.\n", randomNum, attemptsValue)              //Finish game with score
}

func checkNum(computerNum int) int {
	var userEntry string
	var attempts int
	fmt.Scanln(&userEntry)                  //Scans for userinput
	userInt, err := strconv.Atoi(userEntry) //converts to int
	if err != nil {                         //error check
		panic(err)
	}
	for userInt != computerNum { //if user int is not equal to the random number generated, it loops it back again
		attempts++                 //every attempt adds to this
		if userInt > computerNum { //if number is too big
			fmt.Println("Too big! Try again: ")
		}
		if userInt < computerNum { //too small
			fmt.Println("Too small! Try again: ")
		}
		fmt.Scanln(&userEntry)                 //Gets user input and assigns it to our string
		userInt, err = strconv.Atoi(userEntry) //converts it into an int
		if err != nil {
			panic(err)
		}
	}
	return attempts
}
