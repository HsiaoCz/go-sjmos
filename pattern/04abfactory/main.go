package main

import "fmt"

// 抽象工厂方法模式

// --抽象层---
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("china apple..")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("china banana..")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("china pear..")
}

// 中国工厂类
type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	apple := new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	banana := new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	pear := new(ChinaPear)
	return pear
}

// 日本产品类

// 逻辑层
func main() {
	// 需求1 需要中国的苹果，香蕉，梨
	// 创建一个中国工厂
	// var cf AbstractFactory
	cf := new(ChinaFactory)

	// 生成苹果 香蕉 梨
	// var cApple AbstractApple
	cApple := cf.CreateApple()
	cApple.ShowApple()

	cbanana := cf.CreateBanana()
	cbanana.ShowBanana()

	cpear := cf.CreatePear()
	cpear.ShowPear()
}

// 这里面省略了一部分
