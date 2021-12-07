package main

import (
	"fmt"
)

const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int = dogSpecialist
	fmt.Printf("%v", dogSpecialist == specialistType)
}
