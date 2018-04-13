/*
指针：
	go虽然保留了指针，但与其他编程语言不同的是，在go当中不支持
指针运算以及“->”运算符，而是直接采用“.”选择符来操作指针目标对象
的成员
--操作符“&”取变量地址，使用“*”通过指针 间接 访问目标对象
--默认值为nil而非NULL
 */
package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a // 定义指向int a的指针p
	//var p = &a
	fmt.Println(p) // 直接输出p是p所指向的内存地址
	fmt.Println(*p) //  访问指针指向地址的值
}

