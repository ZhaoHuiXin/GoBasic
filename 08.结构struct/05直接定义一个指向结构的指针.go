/*
结构是值类型，对它进行传递的时候也是值拷贝
 */
 package main

import "fmt"

type person5 struct {
 	Name string
 	Age int
 }
func main() {
	// 直接定义一个指向结构的指针
	a := &person5{
		Name:"love",
	}
	// 当a是指针时，对属性操作的时候无需使用*a.属性，直接a.属性即可
	fmt.Println(a)
	A5(a)
	a.Name = "boy"
	fmt.Println(a)
}

func A5(per *person5){
	per.Age = 18
	per.Name = "girl"
	fmt.Println("A5",per)
}