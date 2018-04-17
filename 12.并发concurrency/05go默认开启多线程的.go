package main

import(
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c5 := make(chan bool)
	for i := 0;i<10;i++{
		go Go5(c5,i)
	}
		<-c5
		//time.Sleep(2 *time.Second)
}

func Go5(c chan bool,index int){
	a := 1
	for n := 0;n<1000000;n++{
		a += n
	}
	fmt.Println(index, a)

	if index==9{
		c <- true
	}
}
