package main

import "fmt"

type DIYINT int

func main() {
	var a DIYINT
	a = 12
	a.Increase()
	fmt.Println(a)
	a.Increase()
	fmt.Println(a)
	a.Increase()
	fmt.Println(a)
	a.Increase()
	fmt.Println(a)
}
func (reca *DIYINT)Increase(){
	*reca = *reca + 100
}