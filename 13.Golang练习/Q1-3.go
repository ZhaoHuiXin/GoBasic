package main

import "fmt"

func main() {
	//a := "A"
	//for i:=1;i<=100;i++{
	//
	//	fmt.Println(i,a)
	//	a += "A"
	//}
	//

	//b := "asSASA ddd dsjkdsjs dk"
	//i := 0
	//for _,v := range b{
	//	if v != 32{
	//		i ++
	//	}
	//}
	//fmt.Println(i)

	//c := "asSASA ddd dsjkdsjs dk"
	//k := []rune(c)
	//k[0] = 'a'
	//k[1] = 'b'
	//k[2] = 'c'
	//fmt.Println(string(k))

	//str := "dsjkdshdjsdh....js"
	//fmt.Printf("String: %s \nLength: %d \nRunes: %d\n",str,
	//	len([]byte(str)),utf8.RuneCount([]byte(str)))

	//s := "123456789"
	//r := []rune(s)
	//copy(r[4:4+3],[]rune("abc"))
	//fmt.Printf("before:%s\n", s)
	//fmt.Printf("After :%s\n",string(r))

	str := "abcdefghijklmn"
	a := []rune(str)
	for i, j := 0, len(a)-1; i<j; i,j=i+1,j-1{
		a[i],a[j] = a[j],a[i]
	}
	fmt.Printf("%s\n",string(a))
}