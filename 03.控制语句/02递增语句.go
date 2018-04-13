/*
	在go当中，++与--是作为语句而不是作为表达式
：表达式是可以放在的等号右边的
 */
package main

import (
	"fmt"
)

func main() {
	a := 1
	a++   // 这种形式是语句
	// a = a++ 这是表达式

	fmt.Println(a) //  访问指针指向地址的值
}
