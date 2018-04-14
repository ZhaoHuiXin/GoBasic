/*
1.slice本身并不是数组，它指向底层的数组
2.作为变长数组的替代方案。可以关联底层数组的局部或者全部
3.slice为引用类型
4.可以直接创建或从底层数组中获取生成slice
5.使用len()获取元素个数，cap()获取slice当前最大容量
6.如果多个slice指向相同底层数组，其中一个的值改变会影响全部
创建slice：
7.make([]T,len,cap) T是type
8.其中cap可以省略，则和len的值相同
一般使用make()创建而不是new，new得到的是指向数组的指针
 */
 package main

import "fmt"

// 一，声明slice
//func main() {
//	var s1 []int // 这是声明slice，而不是数组
//	fmt.Println(s1)
//}

// 二，make声明
func main() {
	s2 := make([]int,3,10) // 内存达到极限会重新分配,内存地址改变
	fmt.Println(len(s2), cap(s2))

}
