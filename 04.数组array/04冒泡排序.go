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

// 一,冒泡排序
func main() {
	a := [...]int{1,100,55,66,41,56,78,99,20}
	fmt.Println(a)

	num:=len(a)
	for i := 0;i<num;i++{
		for j:=i+1;j<num;j++{
			if a[j] < a[i]{
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
	fmt.Println(a)
}






