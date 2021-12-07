package main

import (
	"fmt"
)

func main() {
	var i float64 = 32.475
	a := 10.2
	b := 3.7

	fmt.Printf("%.2f\n", i)

	fmt.Printf("%v, %T", i, i)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	var n complex64 = 3.3 + 2i
	var m complex64 = 2 + 2.5i
	fmt.Printf("%v, %T", n, n)
	fmt.Println(n + m)
	fmt.Println(n - m)
	fmt.Println(n * m)
	fmt.Println(n / m)
	fmt.Println(real(n), imag(n))
}
