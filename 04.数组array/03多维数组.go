/*数组array
1.定义数组的格式:var<varName>[n]<type>, n>=0
2.注意,数组的长度也是类型的一部分,因此具有不同长度的数组为不同类型
3.注意区分指向数组的指针和指针数组
4.数组在Go中为值类型
5.数组之间可以使用==或!=进行比较,但不可以使用 < 或 >
6.Go支持多维数组
 */
 package main

import "fmt"

// 一,定义多维数组
//func main() {
//	a3 := [2][3]int{
//		{1,2,3},
//		{4,5,6}, //这要不加“,”，就得把“}”提上来
//	}
//	fmt.Println(a3)
//}

// 二，同样可以在定义数组时指定索引值的元素
//func main() {
//	a3 := [2][3]int{
//		{1:2},
//		{2:100}, //这要不加“,”，就得把“}”提上来
//	}
//	fmt.Println(a3)
//}

// 三，只有最外层数组可以使用[...]表示元素个数未知
func main() {
	a3 := [...][3]int{
		{1:2},
		{2:100},
		{2:100},
		{2:100},
		{0:120},
	}
	fmt.Println(a3)
}




