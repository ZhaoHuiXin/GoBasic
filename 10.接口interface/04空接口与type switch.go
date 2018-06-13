// type switch 多类型判断
package main

import "fmt"

type empty interface {
	// 空接口中什么都没有，所以go中的所有类型都实现了空接口
}
// 接口USB3
type USB4 interface{
	Name() string // 有返回类型string
	Connecter4
}

// 接口Connecter4
type Connecter4 interface{
	Connect()

}

type PhoneConnecter4 struct {
	name string
}

func (reca PhoneConnecter4)Name()string{
	return reca.name
}


func (reca PhoneConnecter4)Connect(){
	fmt.Println("Connect:",reca.name)
}

func main() {
	var a USB4
	a=PhoneConnecter4{"hahaha"}
	a.Connect()
	Type_judge(a)
}
// 使用type switch进行多类型判断
func Type_judge(a_type interface{}) {
	// a_type.(type) 这样写让系统去猜它是什么类型
	switch v := a_type.(type){
	case PhoneConnecter4:
		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}
