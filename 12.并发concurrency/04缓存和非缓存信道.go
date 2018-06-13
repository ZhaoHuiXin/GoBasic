package main
import(
	"fmt"
)

// 有缓存的时候，程序会异步照常执行，不管你信道的状态
//var c4 = make(chan bool, 1)

// 无缓存的时候程序会发生同步阻塞
var c4 = make(chan bool)
func main() {
	go Go4()
	fmt.Println("333")
	<- c4 // 333 111 222
	//c4 <- true  // 333 111 222
}
// 都是主进程走得太快
func Go4(){
	fmt.Println("111")
	c4 <- true  // 333 111 222
	//<- c4  // 333 111 222
	fmt.Println("222")

}