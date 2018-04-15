/*
1.反射可以大大提高程序的灵活性，使得interface有更大的发挥余地
2.反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
3.反射会将匿名字段作为独立字段（匿名字段本质）
4.想要利用反射修改对象状态前提是interface.data是settable，即pointer-interface
5.通过反射，可以动态调用方法

 */
package main
import(
	"fmt"
	"reflect"
)

type User struct{
	Name string
	Id int
	Age int
}

func (u User)Hello(){
	fmt.Println("hello world!")
}

func main() {
	u1 := User{Name:"lucky",Id:12356,Age:18}
	Info(u1)
}

func Info(o interface{}){
	t := reflect.TypeOf(o) // 获得类型信息
	fmt.Printf("%T\n",t)
	fmt.Println("type:",t.Name())

	v := reflect.ValueOf(o) // 获得字段信息
	fmt.Printf("%T\n",v)
	fmt.Println(t.NumField(),v.NumField())
	fmt.Printf("%T,%T\n",t.NumField(),v.NumField())
	// 获得字段信息和类型信息
	for i:=0;i<v.NumField();i++{
		f1 := t.Field(i) // 得到的是接口
		f2 := v.Field(i).Interface() // 得到的是值
		//fmt.Println(f1,"----",f2)
		// 打印 字段名，字段类型，字段值
		fmt.Printf("%6s: %v = %v\n",f1.Name,f1.Type,f2)
	}
	// 获得对象方法信息
	for i :=0;i<t.NumMethod();i++{
		m := t.Method(i)
		fmt.Printf("%6s,%v\n",m.Name,m.Type)
	}

}