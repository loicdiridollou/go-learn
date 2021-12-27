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

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		episodes:  []string{},
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
}
