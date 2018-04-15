/* function 函数
闭包
*/
package main

import "fmt"

func main() {
	a := closeA(100)
	fmt.Println(a(50))
	fmt.Println(a(100))
}

func closeA(x int)(func(int) int){
	fmt.Printf("%p\n",&x)
	return func(y int) int{
		fmt.Printf("%p\n",&x)
		return x+y
	}
}
// 上面外层传入的参数x并不是值的拷贝（在一般函数中传入的是int string的值的拷贝）
// 而是指向x的原始地址

