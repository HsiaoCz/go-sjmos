package main

import "fmt"

//	---抽象层----
//
// ----抽象的水果类
type Fruit interface {
	Show()
}

// ----抽象工厂类---
type AbstractFactory interface {
	CreateFruit() Fruit
}

// ----基础模块层-----
type Apple struct{ Fruit }

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{ Fruit }

func (b *Banana) Show() { fmt.Println("this is banana") }

type Pear struct{ Fruit }

func (p *Pear) Show() { fmt.Println("this is pear") }

// -----基础工厂层-----
type AppleFactry struct {
	AbstractFactory
}

func (fac *AppleFactry) CreateFruit() Fruit {
	fruit := new(Apple)
	return fruit
}

type BananaFactry struct {
	AbstractFactory
}

func (fac *BananaFactry) CreateFruit() Fruit {
	fruit := new(Banana)
	return fruit
}

type PearFactry struct {
	AbstractFactory
}

func (fac *PearFactry) CreateFruit() Fruit {
	fruit := new(Pear)
	return fruit
}

// 业务逻辑层

func main() {
	// 需求1：需要一个具体的苹果对象
	// 1-先要一个具体的苹果工厂
	appleFactory := new(AppleFactry)
	apple := appleFactory.CreateFruit()
	apple.Show()

	// 需要一个香蕉的对象
	bananaFactory := new(BananaFactry)
	banana := bananaFactory.CreateFruit()
	banana.Show()

	// 需要一个梨
	pearFactory := new(PearFactry)
	pear := pearFactory.CreateFruit()
	pear.Show()

}
