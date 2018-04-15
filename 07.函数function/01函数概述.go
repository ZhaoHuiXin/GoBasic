/* function 函数
1.Go函数不支持 嵌套，重载，默认参数
2.GO函数支持：
	无需声明原型、不定长度变参、多返回值、命名返回值参数
	匿名函数、闭包
3.定义函数使用关键字func,且做大括号不能另起一行
4.函数也可以作为一种类型使用
*/
package main

import "fmt"

// A()参数列表，(参数名称1 参数类型,参数名称2 参数类型)
// A()参数列表，(参数名称1,参数名称2,参数名称3 相同的参数类型)
// A(...) 单个返回值类型[(返回值1类型，返回值2类型)]，没有返回值可不写
/*
A()(a,b,c int){....  可以的,这是 命名返回值，命名返回值，在
函数体内return 的时候可以不用return a,b,c ；只要函数体内存在局部变量
a，b，c 就能被返回;
另外，多个返回值是同意类型时，命名返回值起到简化作用，
否则必须写为A(...)(int,int,int),要不然只认为有一个返回值
*/

// 1.无命名返回值
//func A() (int,int,int){
//	a,b,c := 1,2,3
//	return a,b,c
//}

// 2.命名返回值
//func A() (a,b,c int){
//	a,b,c = 1,2,3  // 若是命名了返回值，函数体内的:都省略了
//	return
//}

// 3.不定长变参,必须作为参数列表中的最后一个参数
//func main() {
//	A("name",2,3,4,5,6,7)
//}
//func A(b string, a ...int){
//	fmt.Println(a)
//}

/*
4.不定长变参是引用slice类型，如果传入的是不可变类型参数，
 那么在函数体内是无法改变传入的参数的；
但是如果传入的是可变参数，在函数体内是可以改变该参数的
*/

// 传入不可变类型;如int型，string型
// 只是参数值的拷贝
func main() {
	a := 1
	b := 2
	fmt.Printf("%p\n",&a)
	fmt.Printf("%p\n",&b)
	fmt.Println(a,b)
	A(a,b)
}
func A(a ...int){
	fmt.Printf("%p\n",&a[0])
	fmt.Printf("%p\n",&a[1])
	fmt.Println(a[0],a[1])
	a[0] = 111
	a[1] = 222
	fmt.Printf("%p\n",&a[0])
	fmt.Printf("%p\n",&a[1])
	fmt.Println(a[0],a[1])
}

// 传入可变类型参数,slice,
// 并不是传一个指针，而是对slice的内存地址进行拷贝
//func main() {
//	a := []int{1,2,3,4,5,6}
//	fmt.Println(a)
//	A(a)
//}
//func A(a []int){
//	a[0][0] = 111
//	a[0][1] = 222
//	fmt.Println(a[0])
//}

// 传入可变类型参数,array
//func main() {
//	a := [6]int{1,2,3,4,5,6}
//	fmt.Println(a)
//	A(a)
//}
//func A(a [6]int){
//	a[0][0] = 111
//	a[0][1] = 222
//	fmt.Println(a[0])
//}

// 如果想改变传入的int和string类型的值，传入指向地址的指针
//func main() {
//	a := 1
//	fmt.Println(a)
//	A(&a)
//}
//func A(a *int){
//	*a = 2
//	fmt.Println(*a)
//}




