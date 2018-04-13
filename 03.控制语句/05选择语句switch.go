/*
1.可以使用任何类型或表达式作为条件语句
2.不需要写break，一旦条件符合自动终止
3.如希望继续执行下一个case，需要使用fallthrough语句
4.支持一个初始化表达式（可以是并行方式），右侧需要跟分号
5.左大括号必须和条件语句在同一行
6.！！在if、for、switch初始化语句中定义的变量，都是局部变量
 */
package main

import "fmt"

// 一
//func main() {
//	a5 := 1
//	switch a5 { //switch后加变量，在内部直接写变量可能等于的值
//	case 0:
//		fmt.Println("a5 = 0")
//	case 1:
//		fmt.Println("a5 = 1")
//	}
//	fmt.Println(a5)
//}

// 二，如果switch后不加表达式，那么要在case后面加表达式
//func main() {
//	a5 := 1
//	switch {
//	case a5 >= 0:
//		fmt.Println("a5 = 0")
//		fallthrough
//	case a5 >= 3: // ！！ fallthrough会忽略条件判断，直接执行该case内的语句
//		fmt.Println("a5 = 1")
//	}
//	fmt.Println(a5)
//}

// 三，switch支持初始化表达式，记得右侧需要跟分号;

func main() {
	switch a5:=1; { // 此时变量a5 只在switch语句块内有效
	case a5>=0:
		fmt.Println("a5 = 0")
		//fallthrough
	case a5<0:
		fmt.Println("a5 = 1")
	}
	//fmt.Println(a5)  // 打印不出来了，a5只在switch语句块内有效
}
