package main

import "fmt"

// 接口USB3
type USB3 interface{
	Name() string // 有返回类型string
	Connecter3
}

// 接口Connecter3
type Connecter3 interface{
	Connect()

}

type PhoneConnecter3 struct {
	name string
}

func (reca PhoneConnecter3)Name()string{
	return reca.name
}


func (reca PhoneConnecter3)Connect(){
	fmt.Println("Connect:",reca.name)
}

func main() {
	var a USB3
	a=PhoneConnecter3{"hahaha"}
	a.Connect()
	Disconnect3(a)
}

// 类型判断,ok pattern模式
// usb.(PhoneConnecter3)括号内放入要判断的类型,
// 判断usb是不是PhoneConnecter3类型
// 这样还不够完善，因为不属于USB3类型的参数传不进来，也就是说
// 不会判断出Unknown decive
// 如果想让所有类型都能传进来，需要将USB3换为空接口  interface{}
// interface就像是python中的object，所有类型的dad
func Disconnect3(usb USB3){
	if pc, ok := usb.(PhoneConnecter3);ok{
		fmt.Println(ok)
		fmt.Println(pc) // pc是main.PhoneConnecter
		fmt.Printf("%T,%T\n",pc,ok)

		fmt.Println("Disconnected:",pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}


