//select默认是阻塞的，只有当监听的channel中
// 有发送或接收可以进行时才会运行，
// 当多个channel都准备好的时候，select是随机的选择一个执行的。
package main

import(
	"fmt"
)

func fibonacci(c, quit chan int){
	x, y := 1, 1
	//for {
	//	select {
	//		case c <- x:
	//		x, y = y, x+y
	//		case v:=<- quit:
	//			fmt.Println("quit",v)
	//			return
	//	}
		for i :=0; i < 10; i++{
			c <- x
			x, y = y, x+y
		}
		<- quit
		fmt.Println("quit",)
		fmt.Println("2",)

		return
}

func main() {
	quit, c := make(chan int ), make(chan int )
	go func(){
		for i:=0; i<10; i++{
			fmt.Println("1",<-c)
		}
		quit <- 999
	}()
	fibonacci(c, quit)
	//fmt.Println(<- quit )
}