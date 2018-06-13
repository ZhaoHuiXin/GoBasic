/*
字段首字母大小写的区别产生作用是以package为级别的，
在一个package中没有区别，都是公有的
 */
package main

import(
	"fmt"
)

type A5 struct {
	name string
}
func main(){
	//var a A5 // 与a := A5{}一样，声明a是A5的实例
	a := A5{}
	a.Print()
	a.name="456"
	fmt.Print(a.name) // main也是可以访问结构A5中的私有字段的
}

// struct的方法是可以访问结构的私有字段(属性)(小写的private)的
func (reca *A5)Print(){
	reca.name = "123"
	fmt.Println(reca.name)
}