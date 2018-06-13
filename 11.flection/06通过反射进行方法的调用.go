/*
3.反射会将匿名字段作为独立字段（匿名字段本质）
4.想要利用反射修改对象状态前提是interface.data是settable，即pointer-interface
5.通过反射，可以动态调用方法

 */
package main

import (
	"fmt"
	"reflect"
)

type User6 struct {
	Id int
	Name string
	Age int
}

func (u User6) Hello(name string,age int){
	fmt.Println("hello",name,", my name is ",u.Name,",age:",age)
}

func main() {
	u := User6{1,"OK",123}
	u.Hello("leo",12)

	// 这里只是简单的演示，真正运用要判断被传入参数是否能被修改
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	// 参数要求传入一个slice，
	// reflect.Value类似于空接口，在类型的上层
	args := []reflect.Value{reflect.ValueOf("jerry"),reflect.ValueOf(999)}
	mv.Call(args)

}