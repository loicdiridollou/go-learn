package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
