/*
结构是一种类型，在相同类型间，变量可以进行相互赋值
 */
 package main

import "fmt"

type person9 struct {
	Name string
	Age int
}
type person10 struct {
	Name string
	Age int
}

// 赋值操作
//func main() {
//	a := person9{"lovv",18}
//	b := person9{}
//	b = a
//	fmt.Println(b)
//}

// 比较操作
func main() {
	a := person9{Name:"li"}
	//b := person10{Name:"li"} // person9和10名称不同所以是不同类型，没有可比性
	b := person9{}
	fmt.Println(b == a)

}