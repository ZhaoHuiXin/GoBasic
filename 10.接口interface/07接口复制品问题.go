package main

import "fmt"

type USB7 interface{
	Name() string // 有返回类型string
	Connecter7
}
type Connecter7 interface{
	Connect()
}
type PhoneConnecter7 struct {
	name string
}
func (reca PhoneConnecter7)Name()string{
	return reca.name
}
func (reca PhoneConnecter7)Connect(){
	fmt.Println("Connect:",reca.name)
}

func main() {
	pc := PhoneConnecter7{"Motor"}
	var a Connecter7
	a = Connecter7(pc)
	a.Connect()

	pc.name = "Nokia"
	// 打印结果还是Motor，说明传入接口的是struct P7 的复制品
	a.Connect()
}

