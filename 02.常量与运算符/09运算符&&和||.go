// go中所有的运算符都是从左到右结合的
//    &&与   ||或
package main

import "fmt"

func main() {
	var a = 0
	if a > 0 && (10/a) > 1{  //如果 &&左边的是false，那么根本不执行右边的语句，所以不会报错
		fmt.Println("ok")
	}
}