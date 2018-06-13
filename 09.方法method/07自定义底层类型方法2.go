package main

import "fmt"

type DIYIN int

func main() {
	var a DIYIN
	a = 12
	a.Increase(88)
	fmt.Println(a)

}
func (reca *DIYIN)Increase(num int){
	// num是int类型，与DIYIN不是同一类型，要进行运算得先转换为同一类型
	*reca += DIYIN(num)
}