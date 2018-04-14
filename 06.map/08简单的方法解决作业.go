// 利用for range部分的知识，将map[int]string的键和值进行交换
package main

import "fmt"

func main() {
	m1 := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	m2 := make(map[string]int,5)
	for k,v := range m1{
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
