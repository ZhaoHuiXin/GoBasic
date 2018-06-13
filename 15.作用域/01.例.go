package main

import "fmt"

var a int
func main() {
	a = 5
	fmt.Println(a)
	f()
}
func f(){
	a :=9
	fmt.Println(a)
	g()
}

func g(){
	fmt.Println(a)
}

