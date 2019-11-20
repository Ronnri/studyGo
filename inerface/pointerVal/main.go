package main

import "fmt"

//值 接收者 & 指针接收者  接口

type animal interface {
	mover
	sayer
}

type mover interface {
	move()
	move2()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

//使用值接收者: 类型的值 和类型的指针 都能够保存到接口变量中
func (p person) move() {
	fmt.Printf("%s在跑\n", p.name)
}

//使用指针接收者实现接口 : 只有类型指针可以保存到接口变量中
func (p *person) move2() {
	fmt.Printf("%s在跑~~~~\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在bb\n", p.name)
}

func main() {
	var m mover
	p1 := &person{
		name: "Zhangsan",
		age:  18,
	}
	//p2是person类型的指针
	p2 := &person{
		name: "Lisi",
		age:  20,
	}
	m = p1
	m.move()
	m = p2
	m.move2()
	var s sayer = p2
	s.say()

	var a animal = p2
	a.move()
	a.move2()
	a.say()

}
