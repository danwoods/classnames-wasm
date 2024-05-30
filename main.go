package main
import "fmt"


func main() { 
// obj := map[string]bool{"name": true} 
  // classnames("hello", "Change", "the ", "World using Golang", obj)
}

func classnames(messages ...interface{}) {
  for _, i := range messages {
    fmt.Println(i)
  }
}

