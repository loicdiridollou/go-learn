package main

import "fmt"

func main() {
	d := sum(1, 2, 3, 4, 5)
	fmt.Println(d)
	dd, err := divide(5., 0.)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dd)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}
