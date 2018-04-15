/*
结构是值类型，对它进行传递的时候也是值拷贝
 */
 package main

import "fmt"

type person4 struct {
 	Name string
 	Age int
 }
func main() {
	a := person4{
		Name:"love",
	}
	fmt.Println(a)
	A4(&a)  // 传入a的地址指针，而不是a的值
	fmt.Println(a)
}

func A4(per *person4){
	per.Age = 18
	per.Name = "girl"
	fmt.Println("A4",per)
}