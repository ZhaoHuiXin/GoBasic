package main

import "fmt"

type(
	byte int8   // byte为int8的别名 可以相互进行转换
	rune int32 // rune为int32的别名 可以相互进行转换
	文本  string //不能叫做别名，只是底层数据结构相同，称为自定义类型
				// `文本`和string仍需显式进行类型转换
	Bytesize int64
)
func main() {
	var b 文本
	b = "中文类型别名"
	fmt.Println(b)
}