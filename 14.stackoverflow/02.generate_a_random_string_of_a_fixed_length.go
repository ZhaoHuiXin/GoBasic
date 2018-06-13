package main

import (
	"math/rand"
	"time"
	"fmt"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz" +
						"ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string{
	b := make([]rune, n)
	for i := range b{
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	result := RandStringRunes(100)
	fmt.Println(result)
	fmt.Println(rand.Intn(52))
}
