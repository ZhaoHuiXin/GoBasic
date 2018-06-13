// go中所有的运算符都是从左到右结合的
package main

import (
	"fmt"
)

func main() {
	fmt.Println(^2)  // 一元运算符；优先级最高
	fmt.Println(1^2)  // 二元运算符
}