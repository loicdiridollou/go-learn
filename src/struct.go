package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func test() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		episodes:  []string{},
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	anotherDoctor := &aDoctor
	anotherDoctor.actorName = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}

func main() {
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b.Name)
}
