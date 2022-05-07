package main

import (
	"fmt"
	"math"
)

func other() {
	number := 50
	guess := 30

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be smaller than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low!")
		}
		if guess > number {
			fmt.Println("Too high!")
		}
		if guess == number {
			fmt.Println("Got it!")
		}
	}
}

func sqrt() {
	myNum := 0.123
	if math.Abs(myNum-math.Pow(math.Sqrt(myNum), 2)) < 0.00001 {
		fmt.Println("Results are the same!")
	} else {
		fmt.Println("Results are different!")
	}
}

func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("one")
		fallthrough
	case i <= 20:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

}
