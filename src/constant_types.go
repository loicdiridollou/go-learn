package main

import (
	"fmt"
)

const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota // ignore the first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// var specialistType int = dogSpecialist
	// fmt.Printf("%v", dogSpecialist == specialistType)
	fileSize := 4000000000.
	fmt.Printf("%.2f", fileSize)

}
