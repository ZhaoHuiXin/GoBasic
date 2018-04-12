// go中所有的运算符都是从左到右结合的
//    &&与   ||或
package main

import "fmt"

const(
	B float64 = 1<<(iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
func main() {
fmt.Println(YB)
	}
