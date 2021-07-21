package main

import "fmt"

func main() {
	switch false{ // <=> switch true {
	case true:
		fmt.Println("hello")
	default:
		fmt.Println("bye")
	}
}
