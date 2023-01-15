package main

import (
	"fmt"
)

func main() {
	// 初始化对象的时候可以使用键值对形式，直接写值的集合，每个值后面都需要添加都逗号
	p1 := person{
		"123",
		"张三",
		18,
	}
	fmt.Println(p1)
	changePerson(&p1)
	fmt.Println(p1)
}

// 结构体，类似Java对象，但是结构体是一个常量，如果函数传值不能修改结构体的内容，如果要修改需要使用结构体的地址进行修改
type person struct {
	id   string
	name string
	age  int8
}

// 传递一个person类的指针
func changePerson(p *person) {
	// 通过*p取p指针中的值
	// 通过指针修改结构体中值的方式有两种，一个是直接取值*p，一个是通过.来修改其中变量
	*p = person{
		"666",
		"老六",
		10,
	}
	p.name = "老八"
}
