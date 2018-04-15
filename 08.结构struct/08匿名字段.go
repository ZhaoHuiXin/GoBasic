/*
结构是值类型，对它进行传递的时候也是值拷贝
 */
 package main

import "fmt"

type person8 struct {
	string
	int
}

func main() {
	a := person8{"lovv",18}
	// a := person8{18,"love"}报异常，
	// 使用匿名字段，在进行结构字面值初始化的时候
	// 必须严格按照结构字段的声明顺序
	fmt.Println(a)
}