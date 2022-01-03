package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required:"true" max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	a := Animal{
		Origin: "jfiroegj",
	}
	fmt.Println(a)
	fmt.Println(field)
	fmt.Println(field.Tag)
}
