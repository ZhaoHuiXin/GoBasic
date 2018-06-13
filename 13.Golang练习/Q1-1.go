package main

import "fmt"

func main() {
	// 1
	//for i := 1;i<=10;i++{
	//	fmt.Println(i)
	//}

	// 2
	//i := 0
	//LABEL1:
	//i++
	//if i >10 {
	//	return
	//}
	//fmt.Println(i)
	//goto LABEL1

	//arr := [...]int{1,2,3,4,5,6,7,8,9,10}
	//for _,v := range arr{
	//	fmt.Println(v)
	//}

	arr := [...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("%v\n", arr)
}