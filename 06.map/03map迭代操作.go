package main

import "fmt"

func main() {
	//for k,v := range map{     // 返回键值对，都是map的拷贝，想对map本身操作map[k]=...
	//}

	// 下面这段代码可以证明v是slice sm元素的拷贝
	//sm := make([]map[int]string,5) //slice中的元素是map
	//for _, v := range sm{          // _返回的是索引，丢弃，v返回的是slice中的元素
	//	//v = make(map[int]string,1)  //对slice中的map进行初始化,1表示容量可省略
	//	v = make(map[int]string)
	//	fmt.Println(v)
	//	v[1] = "ok"
	//	fmt.Println(v)
	//}
	//fmt.Println(sm)

	// 如果真的的想修改slice本身，可以这样做
	sm := make([]map[int]string,5)
	for i := range sm{     // i
		sm[i] = make(map[int]string,1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
