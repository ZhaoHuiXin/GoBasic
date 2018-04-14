package main

import(
	"fmt"
)

func main() {
	// 定义map（字典）
	var m1 = map[int]string{}
	var m2 = make(map[int]string)
	m3 := map[int]string{}
	m4 := make(map[int]string)
	fmt.Println(m1,m2,m3,m4)

	// 定义array（数组）
	var a1 = [...]int{}
	var a2 = [10]int{}
	fmt.Println(a1,a2)

	// 定义slice（切片）
	var s1 = []int{1,2,3}
	var s2 = make([]int,3,10)
	fmt.Println(s1,s2)

	// 定义指向数组的指针
	var p1 *[10]int = &a2 // 指向已知数组
	var p2 = new(*[]int) // 指向空数组
	var p3 *[]int = &s1
	fmt.Println(p1,p2,p3)
}
