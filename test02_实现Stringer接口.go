//fmt.Println是我们常用的一个函数，
//但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:
//type Stringer interface {
//	 String() string
//}
//也就是说，任何实现了String方法的类型都能作为参数被fmt.Println调用,
// 也就是说俺可以通过实现String方法来自定义打印内容呢
package main
import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
