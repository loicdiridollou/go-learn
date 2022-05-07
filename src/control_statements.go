package main

import (
	"fmt"
)

func returnTrue() bool {
	fmt.Println("Returned true")
	return true
}

func main() {

	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
	}

	fmt.Println(number <= guess, number >= guess, number != guess)

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
		fmt.Println(ok)
	}
	if true {
		fmt.Println("The test is true.")
	}
}
