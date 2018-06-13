/* method
9.类型别名(自定义类型)不会拥有底层类型所附带的方法
可以使用自定义类型给底层类型添加自定义方法从而实现一些高级的操作
 */
package main

import "fmt"

type TZ int

func main(){
	var a TZ
	a.Print()
}

func (reca *TZ)Print(){
	fmt.Println("TZ")
}


