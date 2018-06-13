package main

import(
	"fmt"
)
// var c chan bool= make(chan bool)
// 简写：
var c3 = make(chan bool)
func main() {
	go Go3()
	// // 直接关闭信道了，信道的内容都没输出
	// time.Sleep(2 * time.Second)
	for v := range c3{
		fmt.Println(v)
	}
}

func Go3(){
	fmt.Println("go go go go!")
	c3 <- true
	// 要在goroutine中关闭信道，在外关goroutine没执行
	close(c3)
}
