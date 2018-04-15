/*
结构是值类型，对它进行传递的时候也是值拷贝
 */
 package main

import "fmt"

type person3 struct {
 	Name string
 	Age int
 }
func main() {
	a := person3{
		Name:"love",
	}
	fmt.Println(a)
	A3(a)  // 传入的a是值拷贝，不会该变原本的a
	fmt.Println(a)
}

func A3(per person3){
	per.Age = 18
	fmt.Println("A3",per)
}