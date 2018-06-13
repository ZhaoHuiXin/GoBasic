/* goto，break，continue
1.三个语法都可以配合标签使用
2.标签名区分大小写，若不使用会造成编译错误
3.break与continue配合标签可用于多层循环的跳出
4.goto是调整执行位置，与其他两个语句配合标签的效果并不相同
 */
package main

import "fmt"

// 一，break配合标签;跳出外层无限循环
//func main() {
//	//LABEL:
//		for{
//			for i:=0;i<10;i++{
//				if i>2{
//					break //LABEL // 若不使用标签,是跳不出外层的无限循环的
//				}else{
//					fmt.Println(i)
//				}
//			}
//		}
//}

// 二，continue配合标签;跳出内层无限循环
//func main() {
//	//LABEL:
//	for i := 0;i<=10;i++{
//		for{
//			fmt.Println(i)
//			continue // LABEL // 若不使用标签,跳不出内层的无限循环
//		}
//	}
//}

// 三，goto调整程序的执行位置
// 尽量把标签放在goto语句之后,避免程序死循环
func main() {
	var i int = 1
	for {
	LABEL1:
		fmt.Println(i,":")
		fmt.Println("执行label1")
		i++
		if i == 10{
			break
		}
		goto LABEL3
	LABEL2:
		fmt.Println("执行label2")
		goto LABEL1
	LABEL3:
		fmt.Println("执行label3")
		goto LABEL2
	}
}
