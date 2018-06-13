//select
//1.可处理一个或多个channel的发送与接收
//2.同时有多个可用的channel时，按随机顺序处理
//3.可用空的select来阻塞main函数
//4.可设置超时
package main

import "fmt"

//select随即向信道写数据
func main() {
	c := make(chan int)
	go func() {
		for v:=range c{
			fmt.Println(v)
		}
	}()
		select {}
	}
// gui程序应用;当执行一个程序后，进入一个事件管理循环，main函数是不能退出的
// 使用select卡死程序，在程序运行过程中发送一些关闭信号来让程序退出