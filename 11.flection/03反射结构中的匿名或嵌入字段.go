/*
3.反射会将匿名字段作为独立字段（匿名字段本质）
4.想要利用反射修改对象状态前提是interface.data是settable，即pointer-interface
5.通过反射，可以动态调用方法

 */
package main

import (
	"reflect"
	"fmt"
)

type User3 struct{
	Id int
	Name string
	Age int
}

type Manager3 struct {
	User3
	title string
}
func main() {
	m := Manager3{User3:User3{1,"okd",12},title:"666"}
	// reflect是以序号的方式取匿名字段的
	t := reflect.TypeOf(m)

	// 序号为0取出Manager的第一个字段User;
	// {User3  main.User3  0 [0] true} bool表示是否为匿名字段，可以用来做判断
	fmt.Printf("%v\n",t.Field(0))

	// 取出Manager中的User中的Id
	fmt.Printf("%v\n",t.FieldByIndex([]int{0,0}))
	// 取出Manager中的User中的Name
	fmt.Printf("%v\n",t.FieldByIndex([]int{0,1}))

	// 序号为1取出Manager第二个字段title;
	// {title main string  16 [1] false}
	fmt.Printf("%v\n",t.Field(1))

}

