package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func old1() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func old2() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking")
}
