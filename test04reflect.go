package main

import (
	"reflect"
	"fmt"
)

var x float64 = 3.4


func main() {
	v := reflect.ValueOf(x)
	fmt.Println("type:",v.Type())
	fmt.Println("kind is float64: ",v.Kind() == reflect.Float64)
	fmt.Printf("value: \n",v.Float())

	p := reflect.ValueOf(&x)
	v = p.Elem()
	v.SetFloat(7.1)
	fmt.Println("type:",v.Type())
	fmt.Println("kind is float64: ",v.Kind() == reflect.Float64)
	fmt.Printf("value: \n",v.Float())
}