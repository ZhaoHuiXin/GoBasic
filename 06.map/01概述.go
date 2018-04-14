/*
1.类似于其他语言中的哈希表或者字典，以key-value形式存储数据
2.key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
3.map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
map定义：
4.make([keyType]valueType,cap)cap表示容量，可省略
5.超出容量时会自动扩容，但尽量给提供一个合理的初始量
6.使用len()获取元素个数
7.键值对不存在时自动添加，使用delete()删除某键值对
8.使用 for range 对map和slice进行迭代操作
 */
package main

import "fmt"

func main() {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int

	numbers = make(map[string]int)
	// 另一种map的声明方式
	//numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3
}
