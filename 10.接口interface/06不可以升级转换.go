// type switch 多类型判断
package main

import "fmt"


// USB5可以转换为Connecter5,反之不行
type USB6 interface{
	Name() string // 有返回类型string
	Connecter5
}

// 接口Connecter5
type Connecter6 interface{
	Connect()
}


//定义一个只实现Connect一个方法的结构
type TVConnecter struct {
	name string
}
func (reca TVConnecter)Connect(){
	fmt.Println("Connect:",reca.name)
}

func main() {
	tv := TVConnecter{"philiper"}
	var a USB
	// a = USB(tv)  // TVConnecter没有实现USB中的Name()，所以不能被转换
	a.Connect()

}

// 使用type switch进行多类型判断
func Type_judge6(a_type interface{}) {
	// a_type.(type) 这样写让系统去猜它是什么类型
	switch v := a_type.(type){
	case PhoneConnecter5:
		fmt.Printf("%T\n",v)

		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}
