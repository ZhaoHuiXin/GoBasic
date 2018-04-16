/*
3.反射会将匿名字段作为独立字段（匿名字段本质）
4.想要利用反射修改对象状态前提是interface.data是settable，即pointer-interface
5.通过反射，可以动态调用方法

 */
package main

import(
	"reflect"
	"fmt"
)

func main() {
	x := 123
	// 当传入的是指针类型
	v := reflect.ValueOf(&x)
	// fmt.Printf("%v\n",v.Field(0)) 只对reflect.struct类型有效
	// 使用v.Elem()取得value；然后可对传入的参数进行操作
	v.Elem().SetInt(999)
	fmt.Println(x)
}


