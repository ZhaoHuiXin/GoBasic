// iota常量计数器，从0开始，组中每定义1个常量自动递增1
// 通过初始化规则与iota可以达到枚举效果
// 每遇到一个const关键字，iota就重置为0
package main
import(
	"fmt"
)
const(
	// a与b都为“A”
	ea = 'A' // "A"输出 A； 'A'输出65 。
			// Go语言的单引号一般用来表示「rune literal」，即——码点字面量。
			//golang里单引号只能有一个字符，如果输出会返回这个字符的ascii码 ，如果想输出为字符需要用string()函数转换一下。
	eb
	ec = iota // c的值是2
	ed // d的值是3
)
// 星期枚举
const (
	// 第一个常量不可省略表达式
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// 常量一般用大写字母命名，但是这样会被外部调用到，可以在常量名前加一个小写字母或者_

func main() {
	fmt.Println(ea,eb,ec,ed)
	fmt.Println(
		Monday,
		Thursday,
		Sunday,
	)
}