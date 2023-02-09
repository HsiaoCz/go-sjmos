package main

import "fmt"

// 为什么需要工厂类

// 假如有一个fruit对象

type Fruit struct {
	//....
	//...
	//...
	//...
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("我是apple")
	}
	if name == "banana" {
		fmt.Println("我是banana")
	}
	if name == "pear" {
		fmt.Println("我是pear")
	}
}

func NewFruit(name string) *Fruit {
	fruit := new(Fruit)

	if name == "apple" {
		// 创建apple的逻辑
		fmt.Println("-----apple-----")
	}
	if name == "banana" {
		// 创建banana的逻辑
		fmt.Println("-----banana-----")
	}
	if name == "pear" {
		// 创建pear的逻辑
		fmt.Println("-----pear-----")
	}
	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.Show("apple")
}

// 如果类的构造方法非常复杂
// 使主函数和类的构造逻辑耦合
// 没有工厂 使业务逻辑和基础类模块耦合
// 工厂模式 使业务逻辑和工厂耦合 工厂和基础模块耦合
