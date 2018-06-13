package main

import "fmt"

func main() {
	// 声明一个key是字符串，值为int的字典,
	var m map[int]map[int]string // 这种方式的声明需要在使用之前使用make初始化
	m = make(map[int]map[int]string)
	a, ok := m[2][1] // 多返回值，ok是布尔类型 告诉我们键值对是否存在，第一次要有:表示创建
	println(a, ok)
	if !ok{
		m[2]=make(map[int]string)  // 不存在使用make进行初始化
	}
	m[2][1] = "good"
	a, ok = m[2][1]  // 再次使用不需要 : 来定义
	fmt.Println(a,ok)

}
