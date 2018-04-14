package main

import "fmt"
// 一、如果多个slice指向相同底层数组，其中一个的值改变会影响全部
//func main() {
//	a := []int{1,2,3,4,5}
//	s1 := a[2:5]
//	s2 := a[1:3]
//	fmt.Println(a,s1,s2)
//	s1[0] = 999
//	fmt.Println(a,s1,s2) // 打印结果说明slice只是指向底层数组，
//}

// 一、如果多个slice指向相同底层数组，其中一个数组因为append而重新分配了内存地址，
// 那么它不会再受之前指向相同底层数组的slice的影响
func main() {
	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(a,s1,s2)
	s2 = append(s2,123,456,789,000,555,44,444,666,66,88,788,9,98,799)
	s1[0] = 9898989
	fmt.Println(a,s1,s2) // 打印结果说明slice只是指向底层数组，
}