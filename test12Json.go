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

	var s map[string]interface{}
	jstr := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
             {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
		 	json.Unmarshal([]byte(jstr), &s)
	fmt.Println(s)
	fmt.Println(s["servers"])
	m, _ := s["servers"].([]map[string]interface{})

	fmt.Println("-------------------------------------------------")
	fmt.Println(m)
	//fmt.Println(m[0])
	//if !ok{
	//	fmt.Println("it's not ok for type string!")
	//	return
	//}
	//fmt.Printf("%T \n",value)
	//fmt.Println(value)
	//fmt.Println(value["servers"])
	//s2 :=value["servers"]
	//m, ok := s2.([]map[string]interface{})
	//fmt.Println(m[1])
	//fmt.Println(m[0]["serverName"])
}

