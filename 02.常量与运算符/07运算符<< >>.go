// go中所有的运算符都是从左到右结合的
package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 10)  // 1MB 10000000000 2的10次方
	fmt.Println(1<<10<<10)  // 1GB 2的20次方
}