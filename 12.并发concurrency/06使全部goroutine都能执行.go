package main

import(
	"fmt"
)

func main() {
	// 不设置runtime.GOMAXPROCS也直接使用多核执行,注意 Go 1.5默认方式
	//runtime.GOMAXPROCS(runtime.NumCPU())
	c6 := make(chan bool,10)
	for i := 0;i<10;i++{
		go Go6(c6,i)
	}
	for i := 0;i<10;i++{
		<-c6
	}
}

func Go6(c chan bool,index int){
	a := 1
	for n := 0;n<1000000;n++{
		a += n
	}
	fmt.Println(index, a)

	c <- true
}
