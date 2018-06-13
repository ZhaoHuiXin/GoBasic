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

type User5 struct {

	Id int
	Name string
	Age int
}
func main() {
	u := User5{1,"ok",18}
	Chge(&u)
	fmt.Println(u)
}
// 要求被修改的对象类型是 reflect.Ptr(指针)，
// 要求接口必须是settable
func Chge(o interface{}){
	v := reflect.ValueOf(o)
	fmt.Printf("%v\n",v.Kind())
	// 判断传入的参数是指针，并且参数的值可以被修改
	if v.Kind() == reflect.Ptr && v.Elem().CanSet(){
		// v取得实际的对象
		v = v.Elem()
	}else{
		fmt.Println("xxx")
		return
	}
	// 利用字段名称取字段值
	f:=v.FieldByName("Name")
	// 判断如果f不合法（invalid），报错
	if !f.IsValid(){
		fmt.Println("Bad")
	}

	if f.Kind()==reflect.String{
		f.SetString("BYB")
	}
}

