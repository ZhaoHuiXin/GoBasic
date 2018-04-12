// 常量的值在编译时就已经确定，程序运行时不可修改
// 常量定义格式与变量基本相同
// =右侧必须是常量或者常量表达式
// 常量表达式中的函数必须是内置函数；
// 因为只有内置函数在编译时就是已经确定了的，自定义函数只有运行时才知道结果
package main

import(
	"fmt"
)

const a int = 1
const b = "A"
const (
	c = a
	d = a + 1
	e = a + 2
	q,w = 1,2
	r,t  // 在定义常量组时，如果不提供初始值，表示将使用上行的表达式
)
const x,y,z = 1,"2",'c'

func main() {
	fmt.Println(r,t)
}