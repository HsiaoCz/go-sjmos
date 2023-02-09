package main

import "fmt"

// Clothes 穿衣服的方式
type Clothes struct {
	// 工作穿衣服
	// 逛街穿衣服
}

// OnWork 工作的装扮
func (c *Clothes) OnWork() {
	fmt.Println("工作的装扮")
}

// OnShop 逛街的装扮
func (c *Clothes) OnShop() {
	fmt.Println("逛街的装扮")
}

func main() {
	c := Clothes{}

	fmt.Println("在工作...")
	//工作的业务
	c.OnWork()

	fmt.Println("在逛街")
	//逛街的业务
	c.OnShop()

	// 工作业务
	fmt.Println("在工作...")
	cw := ClothesWork{}
	cw.Style()

	//逛街的业务
	fmt.Println("在逛街")
	cs := ClothesShop{}
	cs.Style()

	//这样就不容易混淆了
	//我们能很容易的知道方法是属于哪个对象的

}

//这个时候如果我们想让逛街的时候也穿工作的衣服
//有两种方式
//将逛街时装扮改成工作的装扮
//或者直接调用OnWork
//但是这样做就带了问题 容易造成理解上的混淆

//使用单一职责原则

// 拆分成两个类

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}
