/* defer
1.执行方式类似于其他语言中的析构函数，在函数体执行结束后按照
调用顺序的 相反顺序 逐个执行
2.即使函数发生 严重错误 也会执行defer
3.支持匿名函数的调用
4.常用于资源清理、文件关闭、解锁及记录时间等操作
5.通过与匿名函数配合可在return之后 修改 函数计算结果
6.如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时
即已获得了拷贝，否则则是引用某个变量的地址

7.go没有异常机制，但是有panic/recover模式来处理错误
8.panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
package main

import "fmt"

// 一、defer的调用顺序
// a程序执行结束后defer调用c b
//func main() {
//	fmt.Println("a")
//	defer fmt.Println("b")
//	defer fmt.Println("c")
//}

// 执行顺序987...0，此时i作为参数值被拷贝到打印函数fmt.Println中
//func main() {
//	for i:=0;i<10;i++{
//		defer fmt.Println(i)
//	}
//}

// 执行结果987...0，i作为defer时匿名函数的参数，则在定义defer时即已获得了拷贝
//func main() {
//	for i:=0;i<10;i++{
//		defer func(i int) {
//			fmt.Println(i)
//		}(i)  // defer是要执行的，就像a()
//	}
//}

// 执行结果10 10 10，i不是作为defer时匿名函数的参数，只是外层函数的局部变量
// 此时i是作为引用传到defer的匿名函数中的;想引用外层函数的局部变量只能使用匿名函数
// 退出循环体的时候i已经变成了10
func main() {
	for i:=0;i<10;i++{
		defer func() {
			fmt.Println(i)
		}()  // defer是要执行的，就像a()
	}
}



