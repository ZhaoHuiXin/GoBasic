/* method
1.Go中虽没有class，但是依旧存在method
2.通过显式说明receiver来实现与某个类型的组合
3.只能为同一包中的类型定义方法
4.receiver可以时类型的值或者指针;编译器通过接收者的类型来判断method属于哪一结构
5.不存在方法重载
6.可以使用值或指针来调用发那个法，编译器会自动完成转换
7.从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的
第一个参数(Method Values vs.Method Expression)
8.如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
9.类型别名(自定义类型)不会拥有底层类型所附带的方法
10.方法可以调用结构中的非公开字段
 */
 package main

import "fmt"

type A1 struct {
	Name string
 }
type B1 struct {
	Name string
 }

func main() {
	a := A1{}
	a.A_Print()

	b := B1{}
	b.A_Print()
}

// func(a A1)括号中就是一个接收者a,接收者的类型是A1；其实就是函数的第一个参数
// func(a A1)后面是方法的名称
func (a A1) A_Print() {
	fmt.Println("A")
}

// 方法是和结构struct绑定的，不同结构的同名方法不同
// 但是同一结构不能定义两个同名方法
func (b B1) A_Print() {
	fmt.Println("B")
}