package main

import (
	"fmt"
)

func main() {

	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)

	fmt.Println(len(statePopulations))

	sp := statePopulations
	delete(sp, "Texas")
	fmt.Println(statePopulations)
	fmt.Println(sp)
}
