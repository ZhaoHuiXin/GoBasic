// type switch 多类型判断
package main

import "fmt"


// USB5可以转换为Connecter5,反之不行
type USB5 interface{
	Name() string // 有返回类型string
	Connecter5
}

// 接口Connecter5
type Connecter5 interface{
	Connect()
}

type PhoneConnecter5 struct {
	name string
}
func (reca PhoneConnecter5)Name()string{
	return reca.name
}
func (reca PhoneConnecter5)Connect(){
	fmt.Println("Connect:",reca.name)
}

func main() {
	pc := PhoneConnecter5{"Nokia"}
	var a Connecter5
	a=Connecter5(pc)
	fmt.Printf("11,%T\n",a)
	a.Connect()
	//a.Name() // a是Connecter已经不存在Name()方法了
	Type_judge5(a)

}

// 使用type switch进行多类型判断
func Type_judge5(a_type interface{}) {
	// a_type.(type) 这样写让系统去猜它是什么类型
	switch v := a_type.(type){
	case PhoneConnecter5:
		fmt.Printf("%T\n",v)

		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}
