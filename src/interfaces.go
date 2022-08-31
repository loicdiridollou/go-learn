package main

import "fmt"

func main() {
	fmt.Println("Hello go")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello go!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
