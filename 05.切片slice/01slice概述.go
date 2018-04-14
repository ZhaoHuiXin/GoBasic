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

 import(
	 "fmt"
 )

// 一，声明slice
//func main() {
//	var s1 []int // 这是声明slice，而不是数组
//	fmt.Println(s1)
//}

// 二，slice取数组中的元素
func main() {
	var a1 = [10]int{1,2,3,9:100} // 定义一个数组，可以混用定义方式
	//a1 := [10]int{1,2,3}
	fmt.Println(a1)

	var s1 = a1[9] // 取数组中1个元素
	fmt.Println(s1)

	var s2 = a1[5:10] // 取数组中多个元素;包左不包右
	fmt.Println(s2)


}
