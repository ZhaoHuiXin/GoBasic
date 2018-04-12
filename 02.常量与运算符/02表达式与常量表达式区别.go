package main

import(
	"fmt"
)

var ss = "123"  // ss是全局变量，在编译的时候并没有对它进行处理
				// 即在编译的时候，ss是不存在的
const (
	p = len(ss) // 报错：提示常量的初始化必须是常量。
				// 此表达式不是常量表达式
)

func main() {
	fmt.Println(a)
}