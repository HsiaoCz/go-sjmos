package main

import "fmt"

// ---- 抽象层-----
type Fruit interface {
	Show() //接口方法
}

// -----实现层-----
type Apple struct{ Fruit }

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{ Fruit }

func (b *Banana) Show() { fmt.Println("this is banana") }

type Pear struct{ Fruit }

func (p *Pear) Show() { fmt.Println("this is pear") }

// ---工厂层----
type Factory struct{}

func (f *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		// apple构造初始化
		fruit = new(Apple) //满足多态条件的赋值，父类指针指向子类对象
	}
	if kind == "banana" {
		// Banana构造初始化
		fruit = new(Banana) //满足多态条件的赋值，父类指针指向子类对象
	}
	if kind == "pear" {
		// pear构造初始化
		fruit = new(Pear) //满足多态条件的赋值，父类指针指向子类对象
	}
	return fruit
}

func main() {

	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()

}
