//接口调用不会做receiver的自动转换
package main

import "fmt"

type person5 struct {
	Name string
	Age int
}
func A5(per *person5){
	per.Age = 18
	per.Name = "girl"
	fmt.Println("A5",per)
}
func main() {
	// 直接定义一个指向结构的指针
	a := &person5{
		Name:"love",
	}
	// 当a是指针时，对属性操作的时候无需使用*a.属性，直接a.属性即可
	fmt.Println(a)
	A5(a)
	// receiver 自动转换，可以调用非指针方法集，而不用*a.Name
	a.Name = "boy"
	fmt.Println(a)
}



// USB9 接口调用时接口接收的是什么类型就得给它传什么类型
// 指针类型的receiver 的方法集包含了 不是针类型的receiver 传进来的方法集
// 接口传的都是复制品，即值拷贝的方法集，索引接口调用不会进行自动转换；
// 即不不能使用指针的方法集
type USB9 interface{
	Name() string // 有返回类型string
	Connecter9
}
type Connecter9 interface{
	Connect()
}
type PhoneConnecter9 struct {
	name string
}
/*cannot use PhoneConnecter3
literal (type PhoneConnecter3) as
type USB3 in assignment:
*/
func (reca PhoneConnecter9)Name()string{
	return reca.name
}
func (reca PhoneConnecter9)Connect(){
	fmt.Println("Connect:", reca.name)
}

