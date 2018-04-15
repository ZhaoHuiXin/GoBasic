package main

import "fmt"

type DIYI int

func main() {
	var a DIYI
	a = 12
	a.Increase(88)
	fmt.Println(a)

}
func (reca *DIYI)Increase(num DIYI){
	*reca += num
}