/* defer
1.执行方式类似于其他语言中的析构函数，在函数体执行结束后按照
调用顺序的 相反顺序 逐个执行
2.即使函数发生 严重错误 也会执行defer
3.支持匿名函数的调用
4.常用于资源清理、文件关闭、解锁及记录时间等操作
5.通过与匿名函数配合可在return之后 修改 函数计算结果
6.如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时
即已获得了拷贝，否则则是引用某个变量的地址

7.go没有异常机制，但是有panic/recover模式来处理错误
8.panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
package main

import "fmt"

func main() {
	A6()
	B6()  // 当进行B的时候函数就终止了，C是不会执行的
	C6()
}
func A6(){
	fmt.Println("func A6")
}

func B6() {
	//注意defer语句放在panic之前；因为程序panic就不会再往下执行了
	defer func(){
		// check panic是否存在；
		// 正常情况下recover()会返回nil
		// 调用recover会返回panic的信息；如果返回的不是nil说明程序处在恐慌状态会引发panic
		// recover会跳过panic继续使程序执行
		// 如果返回nil，说明什么都不需要做
		// 利用defer中的recover进行恢复;当程序panic后，会先执行最后一个defer
		if err := recover();err != nil{
			fmt.Println("recover inb")

		}
	}()

	//defer func(){
	//	recover()
	//}()
	panic("panic in B6")

}

func C6() {
	fmt.Println("Func C6")
}


