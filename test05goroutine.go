package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	// 在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS
	// 的初始值由1改为了运行环境的CPU核数
	// 如果不手动设为1 那么是并行的，main直接运行完成，goroutine来不及运行;
	// 除非设置等待时间
	runtime.GOMAXPROCS(1)
	go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行
	time.Sleep(1 * time.Second)
}