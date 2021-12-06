package main

import "fmt"

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	sason        int    = 11
)

var (
	counter int = 0
)

func main() {
	// variables
	a := 10
	b := 3

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	a = 8
	fmt.Println(a << b)
	fmt.Println(a >> b)
}
