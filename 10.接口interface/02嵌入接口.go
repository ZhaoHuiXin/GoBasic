package main

import "fmt"

// 接口USB2
type USB2 interface{
	Name() string // 有返回类型string
	Connecter2
}

// 接口Connecter2
type Connecter2 interface{
	Connect()
}

type PhoneConnecter2 struct {
	name string
}

func (reca PhoneConnecter2)Name()string{
	return reca.name
}


func (reca PhoneConnecter2)Connect(){
	fmt.Println("Connect:",reca.name)
}

func main() {
	var a USB2// a的类型是USB
	//PhoneConnecter2实现了USB接口的方法，因此可以赋值给a
	a=PhoneConnecter2{"hahaha"}
	// a.name="hahaha" USB中没有name字段，这样赋值会失败，
	// 但是放到字面值初始化中，它会自动寻找name
	a.Connect()
	Disconnect2(a)
}

/*
要求传进类型为USB的变量，校验
a := PhoneConnecter{"hahaha"}
是否实现了USB接口。
*/
func Disconnect2(usb USB2){
	fmt.Println("Disconnected.")
}


