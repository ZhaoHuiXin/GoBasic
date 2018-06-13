/*
结构是值类型，对它进行传递的时候也是值拷贝
 */
 package main

import "fmt"

type person7 struct {
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}

func main() {
	a := person7{
		Name:"loove",
		Age:18,
	}
	// 匿名结构的初始化只能通过这种方法：
	a.Contact.Phone = "110"
	a.Contact.City = "BeiJing"
	fmt.Println(a)
}