// copy前后不会影响slice的指向的
package main

import "fmt"

// 一、如果多个slice指向相同底层数组，其中一个的值改变会影响全部
func main() {
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{7,8,9}
	fmt.Println(s1,s2) // 将 s2 拷贝到 s1 中
	copy(s1,s2)
	fmt.Println(s1,s2) // [7 8 9 4 5 6] [7 8 9] 从头开始替换

	s3 := []int{11,22,33,44,55,66}
	copy(s2,s3) // s3比s2长，但是也会替换，多了的被舍弃了
	fmt.Println(s2,s3)

	s4 := []int{100,200,300}
	s5 := []int{666,777,888,999}
	copy(s4[1:2], s5[:2]) // 不可以使用索引如s4[2]，但是可以使用切片拷贝
	fmt.Println(s4,s5)
}
