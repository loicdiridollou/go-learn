package main

import (
	"fmt"
)

func main() {

	/*
		var s rune = 'a'
		fmt.Printf("%v, %T\n", s, s)
		//fmt.Printf("%v, %T", string(s[2]), s[2])
	*/

	const myConst float64 = 1.57
	fmt.Printf("%v, %T", myConst, myConst)
	fmt.Println("\n ")
	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T", a+b, a+b)
}
