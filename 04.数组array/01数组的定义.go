/*数组array
1.定义数组的格式:var<varName>[n]<type>, n>=0
2.注意,数组的长度也是类型的一部分,因此具有不同长度的数组为不同类型
3.注意区分指向数组的指针和指针数组
4.数组在Go中为值类型
5.数组之间可以使用==或!=进行比较,但不可以使用<或>
6.Go支持多维数组
 */
 package main

import "fmt"

// 一,不同长度的数组为不同类型
//func main() {
//	var a1 [2]int
//	//var b1 [1]int // 不同类型数组间不能进行赋值操作
//	var b1 [2]int
//	b1 = a1
//	fmt.Println(b1)
//}

// 二,定义数组
func main() {
	a1 := [2]int{1} // 字面值不够,会用零值填充
	fmt.Println(a1)

	b1:=[20]int{19:1} // 使用数组索引给第20个元素赋值
	fmt.Println(b1)

	c1:=[...]int{1,2,3,4,5} // 使用...来定义一个元素值已知的数组
	fmt.Println(len(c1))
	fmt.Println(c1)

	d1 := [...]int{19:1} // 使用索引配合...定义数组,会生成一个符合条件的最低要求的数组
	fmt.Println(d1)

	e1 :=[...]int{19:9} // 索引从0开始
	var p *[20]int = &e1  // 定义指向数组的指针
	fmt.Println(p)  // 打印结果是 取数组的地址

	x1, y1 := 1, 2
	f1 := [2]*int{&x1,&y1} // 定义一个长度为2的数组,保存的元素是两个变量的指针
	fmt.Println(f1)
	fmt.Println(*f1[0]) // 打印数组中第一个指针指向的值
}

