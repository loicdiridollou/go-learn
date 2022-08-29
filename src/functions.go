// package implements some functions items
package main

import (
	"fmt"
)

func main() {
	name := "Stacey"
	greeting := "Hello "
	sayGreeting(&greeting, &name)
	fmt.Println(name)
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of index is ", idx)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}
