// go中所有的运算符都是从左到右结合的
package main

import "fmt"

func main() {
	/*
		6 : 0110
		11: 1011
	-----------------
	&	    0010 =2 与操作符：两个都为1结果为1，有一个不为1结果是0
	|		1111 =15 或操作符：有一个数为1，结果为1
	^		1101 =13 2个中只有1个为1，结果为1
	&^ 		0100 =4 结果来源于第一个数字，第二个数字位上为1，将第一个数对应位置数字强制改为0，其他不变
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6^11)
	fmt.Println(6&^11)
}