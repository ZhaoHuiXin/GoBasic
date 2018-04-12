package main

import "fmt"

type(
	byte int8   // 将int8别名为byte
	rune int32
	文本  string
	Bytesize int64
)
func main() {
	var b 文本
	b = "中文类型别名"
	fmt.Println(b)
}