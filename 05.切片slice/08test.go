package main

import "fmt"

func main() {
		a := [8]int{1,2,3,4,5,6,7,8}
		s1 := a[1:2:6]
		fmt.Println(len(s1),cap(s1))
		s2 := make([]int,1,5)
		s2[0] = 2
		fmt.Println(len(s2),cap(s2))
}
