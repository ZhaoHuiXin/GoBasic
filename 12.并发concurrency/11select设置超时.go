//select设置超时
package main
import(
	"fmt"
	"time"
)

func main() {
	c:=make(chan bool)
	//close(c)
	select{ // 一直运行着的
		case v:= <-c:
			fmt.Println(v)

		case <- time.After(3 * time.Second):
			fmt.Println("Timeout")
	}
	//close(c)
}