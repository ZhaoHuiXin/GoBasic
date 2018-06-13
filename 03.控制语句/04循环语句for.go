/*
1.go只有for一个循环语句关键字，支持3中形式
2.初始化条件和步进表达式可以时多个值
3.条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，
尽量提前计算好条件并以变量或常量代替
4.左大括号必须和条件语句在同一行
 */
package main

import "fmt"

// 形式一,无限循环，应在内部判断break条件（while true)
//func main() {
//	a4 := 1
//	for { //左大括号必须和条件语句或else在同一行
//		a4++
//		if a4 > 3{
//			break
//		}
//		fmt.Println(a4)
//	}
//	fmt.Println(a4)
//}

// 形式二，加判断条件（while i <= 3）
//func main() {
//	a4 := 1
//	for a4 <= 3{
//		a4 ++
//	}
//	fmt.Println(a4)
//}

// 形式三，加判断条件和计数器及步进表达式
func main() {
	a4 := 1
	txt4 := "hello world"
	var n4 int = len(txt4) // 条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数
						   // 尽量提前计算好条件并以变量或常量代替
	for i := 0; i<n4 ; i++{
		a4 ++
	}
	fmt.Println(a4)
}
