//select
//1.可处理一个或多个channel的发送与接收
//2.同时有多个可用的channel时，按随机顺序处理
//3.可用空的select来阻塞main函数
//4.可设置超时
package main
import(
	"fmt"
)
// select接收信道的值
func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool,2)
	go func() {
		for{
			select{
				case v,ok:= <-c1:
					if !ok {
						fmt.Println("c1")
						o <- true
						break
					}
					fmt.Println("c1",v)

				case v,ok := <-c2:
					if !ok{
						fmt.Println("c2")
						o <- true
						break
					}
					fmt.Println("c2",v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	close(c1)
	//close(c2)
	//for i:= 0; i<2;i++{
	//	<- o
	//}
	<- o
}