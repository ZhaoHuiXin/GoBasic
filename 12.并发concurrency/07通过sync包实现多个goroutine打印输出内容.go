package main

import(
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++{
		go Go7(&wg, i)
	}
	wg.Wait()
}

func Go7(wg *sync.WaitGroup, index int){
	a := 1
	for i := 0;i<1000000;i++{
		a += i
	}
	fmt.Println(index,a)
	wg.Done()
}