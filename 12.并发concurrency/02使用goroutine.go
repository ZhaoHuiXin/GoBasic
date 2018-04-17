package main

import(
	"fmt"
)
// var c chan bool= make(chan bool)
// 简写：
var c = make(chan bool)
func main() {

	go Go()
	// 调用sleep，这样才能看见Go执行，否则main执行太快，Go()还没执行
	<- c
}
func Go(){
	fmt.Println("go go go go!")
	c <- true
}
