package main

//import "fmt"

func main() {
	//println(str)
	// obj := map[string]bool{"name": true}
	// classnames("hello", "Change", "the ", "World using Golang", obj)
	// <-make(chan bool)
}

//export classnames
func classnames(messages ...interface{}) {
	println("here")
	println(messages)
	for _, i := range messages {
		println(i)
	}
}
