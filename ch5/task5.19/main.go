package main

import "fmt"

func main() {
	fmt.Println(panicValue())
}

func panicValue() (res interface{}) {
	defer func() {
		if p := recover(); p != nil {
			res = p
		}
	}()
	panic(5)
}
