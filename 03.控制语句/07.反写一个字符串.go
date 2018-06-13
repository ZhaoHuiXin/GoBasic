package main

import "fmt"

func reverse(b string){
	a := []rune(b)
	for i, j := 0, len (a)-1 ; i < j ; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(string(a))
}
func main() {
	reverse("hello")
}

