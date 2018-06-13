package main

import(
	"fmt"
)

func fibonacci(n int, c chan int){
	x, y := 1, 1
	for i :=0; i < n; i++{
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	fmt.Println(cap(c))
	//for i := range c{
	for i := 0; i<11; i++{
		fmt.Println(i)
		v , ok := <-c
		fmt.Println(v, ok)
	}
}
