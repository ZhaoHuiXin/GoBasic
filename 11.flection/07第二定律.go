package main

import "fmt"
func main() {
	func (v Value) Interface() interface{}
	y := v.Interface().(float64)
	fmt.Println(y)
}
