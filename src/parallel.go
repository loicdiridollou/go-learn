package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var msg = "Hello world!"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
