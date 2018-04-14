// copy前后不会影响slice的指向的
package main

import "fmt"

// 一、如果多个slice指向相同底层数组，其中一个的值改变会影响全部
func main() {
	s1 := []int{1,2,3,4,5}
	s2 := s1
	//s2 := s1[0:5] // 和上面等同
	fmt.Println(s2)
}
