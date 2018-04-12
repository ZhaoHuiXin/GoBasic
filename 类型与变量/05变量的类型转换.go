package main

import "fmt"

func main() {
	// go当中不存在隐式转换，所有类型转换必须显式声明(保证类型安全)；
	// 转换只能发生在两种相互兼容的类型之间(int和bool不兼容
	// <value a>[:] = <Type of value a >(<value b>)
	// 当b是全新的类型时有：，因为没有对b进行声明；如果前面已经声明过b，那么不需要：
	var a float32 = 100.1
	fmt.Println(a)
	b := int(a)
	fmt.Println(b)
}