/* function 函数
4.函数也可以作为一种类型使用
类似于python中的函数也是对象
*/
package main

import(
	"fmt"
)

func main() {
	a := fA
	a()
}
func fA(){
	fmt.Println("hello world")
}


