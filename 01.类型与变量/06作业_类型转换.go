package main

import (
	"fmt"
	"strconv"
)
var d int
func main() {
	var a int = 65
	b := string(a)
	//d := 12
	/*string()表示将数据转换成文本格式，因为计算机中存储的任何东西
	本质上都是数字，因此函数自然的认为我们需要的是用数字65表示文本A*/

	// 如果确实要“65”，使用strconv包
	c := strconv.Itoa(a)
	fmt.Println(b,c)
	d,_ = strconv.Atoi(c) // 65,nil
	fmt.Println(d)
}