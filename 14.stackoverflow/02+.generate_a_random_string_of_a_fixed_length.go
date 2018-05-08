package main

import (
	"math/rand"
	"time"
	"fmt"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}
const letterBytes = "abcdefghijklmnopqrstuvwxyz" +
						"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	//result := RandStringRunes(10)
	//fmt.Println(result)
	fmt.Printf("%T\n",string(letterBytes[0]))
	println(string(letterBytes[0]))
	fmt.Printf("%T\n", rand.Int63())
	fmt.Println(rand.Int63())
	var a int64
	a = 1<<63-1
	fmt.Println(a)
}
