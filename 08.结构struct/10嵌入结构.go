/*
结构是一种类型，在相同类型间，变量可以进行相互赋值
 */
 package main

import "fmt"

type human struct {
	Sex int
}

type teacher struct {
	human // 嵌入结构作为匿名字段，系统默认是将结构名称作为字段名称
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}

func main() {
	// 嵌入结构字面值初始化
	a := teacher{Name:"joe",Age:18,human:human{Sex:0}}
	b := student{Name:"joe",Age:18,human:human{Sex:1}}

	// 修改嵌入结构的属性
	a.Name = "lucy"
	a.Age = 100
	//a.human.Sex = 1000
	a.Sex = 999 // 会把嵌入结构的属性都给外层结构
	fmt.Println(a,b)

}