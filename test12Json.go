package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//type Server struct {
	//	ServerName string
	//	ServerIP   string
	//}
	//type Serverslice struct {
	//	Servers []Server
	//}
	//var s Serverslice

	var s interface{}
	jstr := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
             {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
		 	json.Unmarshal([]byte(jstr), &s)
	value, ok := s.(map[string]interface{})
	if !ok{
		fmt.Println("it's not ok for type string!")
		return
	}
	fmt.Printf("%T",value)
	fmt.Println(value)
}