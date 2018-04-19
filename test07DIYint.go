package main

import "fmt"

type myint int

var a myint = 1

func (reca myint)String() string{
	return "it's me "
}

func main() {
	var ta myint = 1
	fmt.Println(ta)
	fmt.Println(a)
}
