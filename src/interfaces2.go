package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("I am an int")
	case string:
		fmt.Println("I am a string")
	default:
		fmt.Println("I don't know")
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
