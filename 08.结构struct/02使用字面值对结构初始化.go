/* go中没有class，所以没有构造函数
但是可以使用字面值对结构进行初始化
 */
 package main

import "fmt"

type person2 struct {
 	Name string
 	Age int
 }
func main() {
	a := person2{
		Name:"love",
	}
	fmt.Println(a)
}