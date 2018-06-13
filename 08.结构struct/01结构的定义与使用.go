/* 结构struct
1.Go中的struct与C中的struct非常相似，GO中没有class
2.使用type<Name>struct{}定义结构，名称遵循可见性规则
3.支持指向自身的指针类型成员
4.支持匿名结构，可用做成员或定义成员变量
5.可以使用字面值对结构进行初始化
6.相同类型的成员可进行直接拷贝赋值
7.支持==与!=比较运算符，但不支持<和>
8.支持匿名字段，本质上是定义了某个类型名为名称的字段
9.嵌入结构作为匿名字段看起来像继承，但不是继承
10.可以使用匿名字段指针
 */

 package main

import "fmt"

// 定义了person结构，
type person struct{
	Name string
	Age int
}

 func main(){
 	a := person{}
 	// 也是使用点操作属性
 	a.Name = "lucky"
 	a.Age = 18
 	fmt.Println(a)
 }
