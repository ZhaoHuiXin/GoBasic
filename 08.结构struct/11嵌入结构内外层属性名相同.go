/*
结构是一种类型，在相同类型间，变量可以进行相互赋值
 */
 package main

import "fmt"

type SA struct {
	Name string
}
type SB struct {
	SA
	Name string
}

func main() {
	a := SB{Name:"abc",SA:SA{Name:"efg"}}
	// a.Name会直接取到结构SB中的Name，它的级别高于嵌入的结构中的Name
	// 如果结构SB中不存在Name字段，那么a.Name调用的是结构SA中的Name
	// a.SA.Name的作用就是当命名冲突时加以区分
	fmt.Println(a.Name,a.SA.Name)
}

// 当嵌入了2个结构，这两个结构有同名字段，那么编译器无法选择用谁的会报错
