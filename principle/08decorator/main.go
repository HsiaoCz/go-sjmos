package main

import "fmt"

// 装饰器模式

// ----抽象层----
type Phone interface {
	Show() // 构建的功能
}

// 抽象的装饰器，装饰器的基础类
// (该类本应该是interface，但是go语言的interface不可以有成员属性)
type Decorator struct {
	phone Phone // 组合进来的抽象的phone
}

func (d *Decorator) Show() {}

// ----具体的构建------
type Huawei struct{}

func (hw *Huawei) Show() { fmt.Println("秀出了huawei手机") }

type Xiaomi struct{}

func (hw *Xiaomi) Show() { fmt.Println("秀出了xiaomi手机") }

// -------具体的装饰器类---------

type MeDecorator struct {
	Decorator //继承基础的装饰器类(主要是继承phone的成员属性)
}

func (md *MeDecorator) Show() {
	md.phone.Show() //调用装饰的构件的原方法
	//-----变成贴膜的手机-----
	fmt.Println("贴膜的手机....")
}

func NewMoDecorator(phone Phone) Phone {
	return &MeDecorator{Decorator{phone: phone}}
}

type KeDecorator struct {
	Decorator //继承基础的装饰器类(主要是继承phone的成员属性)
}

func (ke *KeDecorator) Show() {
	ke.phone.Show() //调用装饰的构件的原方法
	//-----变成带壳的手机-----
	fmt.Println("带壳的手机....")
}

func NewkeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

// 业务逻辑层
func main() {
	huawei := new(Huawei)
	huawei.Show()

	// 用贴膜的装饰器，得到一个新功能的huawei手机
	moHuawei := NewMoDecorator(huawei)
	moHuawei.Show()

	// 壳的装饰器
	keHuawei := NewkeDecorator(huawei)
	keHuawei.Show()
}
