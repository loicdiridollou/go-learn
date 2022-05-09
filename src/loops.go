package main

import (
	"fmt"
)

func old() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}
	fmt.Println(i)
}
