// 利用for range部分的知识，将map[int]string的键和值进行交换
package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	s := make([]int,len(m))  //虽然说map是无序的，s和i的配合取到map中的所有value
	i := 0
	// 通过这个迭代将m中所有的key都存到了slice s当中，但是此时每次slice中的key都是无序的
	for k, _ := range m{
		s[i] = k
		i++
	}
	// 让其间接有序，用到  sort包
	sort.Ints(s) // 对int类型进行排序，默认是升序排列，因为这里map m的可以是int类型。
				// 通过这里也可以看出slice是引用类型的，因为这里并没有返回，
				// 	只是传入s，它就会对slice本身进行操作
	fmt.Println(s)

}
