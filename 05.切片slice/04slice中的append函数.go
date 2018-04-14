/* append
1.可以在slice尾部追加元素
2.可以将一个slice追加在另一个slice的尾部
3.如果最终长度未超过追加到slice的容量，则返回原始slice，即地址不变
4.如果超过追加到的slice的容量则将重新分配slice并拷贝原始数据
 */
package main

import "fmt"

func main() {
	s1 := make([]int,3,6)
	var p1 *[]int = &s1
	fmt.Printf("%p\n",s1) // 注意使用的是Printf 打印的 %p是内存地址
	fmt.Println(p1)
	s1 = append(s1,1,2,3)
	fmt.Printf("%v %p\n",s1,s1) // %v是值，没超过s1的最大容量，地址不变
	s1 = append(s1,4,5,6)
	fmt.Printf("%v %p\n",s1,s1) // 超过了s1的最大容量，重新分配内存，
									  // 并将原始数据拷贝，地址改变
}