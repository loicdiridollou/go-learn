package main

import (
	"fmt"
)

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 93

	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "John"
	fmt.Printf("Student: %v\n", students[1])

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	aa := [...]int{1, 2, 3}
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))

	c := make([]int, 3)
	fmt.Println(c)

	c = append(c, 1)
	fmt.Println(c)

}
