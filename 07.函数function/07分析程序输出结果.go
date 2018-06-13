package main

import "fmt"

func main() {
	var fs = make([]func(),4)

	for i:=0;i<4;i++{
		defer fmt.Println("defer i = ",i)
		defer func(){fmt.Println("defer_closure i = ",i)}()
		fs[i] = func(){fmt.Println("closure i = ",i)}
	}
	for _, f := range fs{
		f()
	}
}