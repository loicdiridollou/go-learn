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

func old_continue() {
	i := 0
	for ; i < 10; i++ {

		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}

func main() {
	s := []int{1, 2, 3}
	for _, v := range s {
		fmt.Println(v)
	}

}
