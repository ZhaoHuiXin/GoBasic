// 利用for range部分的知识，将map[int]string的键和值进行交换
package main

import (
	"sort"
	"fmt"
)

func main() {
	m := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	s1 := make([]int,len(m))
	var i = 0
	for k, _ := range m{
		s1[i] = k
		i ++
	}
	sort.Ints(s1)
	fmt.Println(s1)

	m2 := make(map[string]int,5)
	for v := range s1{
		fmt.Println(s1[v])
		m2[m[s1[v]]] = v
	}
	fmt.Println(m)
	fmt.Println(m2)
}
