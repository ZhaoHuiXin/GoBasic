/*
结构是值类型，对它进行传递的时候也是值拷贝
 */
 package main

import "fmt"

// 匿名结构
//func main() {
//	a := struct{
//		Name string
//		Age int
//	}{
//		Name:"lucy",
//	}
//	fmt.Println(a)
//}

// 匿名结构取指针
func main() {
	a := &struct{
		Name string
		Age int
	}{
		Name:"lucy",
	}
	fmt.Println(a)
}