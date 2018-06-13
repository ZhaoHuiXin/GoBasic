/* 接口interface
1.接口是一个或多个方法签名的集合
2.只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需
显式声明实现了哪个接口，这称为Structural Typing
3.接口只有方法声明，没有实现，没有数据字段
4.将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，
既无法修改复制品的状态，也无法获取指针
5.只有当接口存储的类型和对象都为nil时，接口才等于nil
6.接口调用不会做receiver的自动转换
7.接口同样支持匿名字段方法
8.接口也可以实现类似OOP中的多态
9.空接口可以作为任何类型数据的容器
 */
package main

import "fmt"

type USB interface{
	Name() string // 有返回类型string
	Connect()
}
type PhoneConnecter struct {
	name string
}
func (reca PhoneConnecter)Name()string{
	return reca.name
}
func (reca PhoneConnecter)Connect(){
	fmt.Println("Connect:",reca.name)
}


type PhoneConnecter1 struct {
	name string
}
func (reca PhoneConnecter1)Name()string{
	return  reca.name
}
func (reca PhoneConnecter1)Connect(){
	fmt.Println("Connect:",reca.name+"666")
}

func main() {
	// 第一种调用接口的方式
	var a USB // a的类型是USB
	//PhoneConnecter实现了USB接口的方法Name()，因此可以赋值给a
	a=PhoneConnecter1{"hahaha"}
	// a.name="hahaha" USB中没有name字段，这样赋值会失败，
	// 但是放到字面值初始化中，它会自动寻找name

	// 第二种调用接口的方式
	pc := PhoneConnecter{"++--++"}
	var b USB
	b = USB(pc)
	a.Connect()
	fmt.Printf("%T\n",a)
	Disconnect(a)
	b.Connect()
	fmt.Printf("%T\n",b)
	Disconnect(b)
}

/*
要求传进类型为USB的变量，校验
a := PhoneConnecter{"hahaha"}
是否实现了USB接口。
*/
func Disconnect(usb USB){
	fmt.Println("disconnected.")
}
