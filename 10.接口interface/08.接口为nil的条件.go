package main

import "fmt"

type USB8 interface{
	Name() string // 有返回类型string
	Connecter8
}
type Connecter8 interface{
	Connect()
}
type PhoneConnecter8 struct {
	name string
}
func (reca PhoneConnecter8)Name()string{
	return reca.name
}
func (reca PhoneConnecter8)Connect(){
	fmt.Println("Connect:",reca.name)
}

// 只有当接口存储的类型和对象都为nil时，接口才等于nil
func main() {
	// a中存的就是nil
	var a interface{}
	fmt.Println(a == nil) // true

	var p *int = nil
	// a中存了指向nil的指针，但是a存储的类型是指针
	a = p
	fmt.Println(a == nil) // false
}

