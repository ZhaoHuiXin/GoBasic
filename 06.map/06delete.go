package main

import(
	"fmt"
)

func main() {
	m := map[int]string{10:"a",2:"b"}
	fmt.Println(m)
	delete(m,10)
	fmt.Println(m)
}
