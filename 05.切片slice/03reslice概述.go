/* reslice
1.reslice时索引以被切slice切片为准
2.索引不可以超过被slice的切片的容量cap()值
3.索引越界不会导致底层数组的重新分配而是引发错误
 */
package main

import(
	"fmt"
)

func main() {
	a :=[]string{"a","b","c","d","e","f","g"}
	sa := a[2:5] // slice sa的极限容量是从数组a[2]到尾部，即极限容量为5
	fmt.Println(len(sa),cap(sa))
	fmt.Println(sa)
	sb := sa[1:3]  // reslice 切片sa
	fmt.Println(sb)
	sd := a[3:5]
	fmt.Println(sd)
}