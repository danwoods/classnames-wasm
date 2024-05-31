package main

//import "fmt"

func main() {
	//println(str)
	obj := map[string]bool{"name": true}
	classnames("hello", "Change", "the ", "World using Golang", obj)
}

//export classnames
func classnames(messages ...interface{}) {
	for _, i := range messages {
		println(i)
	}
}
