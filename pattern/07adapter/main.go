package main

import "fmt"

// 适配器模式

// ----适配的目标----
type V5 interface {
	Use5V()
}

// 被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() { fmt.Println("使用220v的电压..") }

// 适配器
type Adopter struct {
	v220 *V220
}

func (a *Adopter) Use5V() {
	fmt.Println("使用适配器进行充电...")
	// 调用适配者的方法
	a.v220.Use220V()
}

func NewAdopter(v220 *V220) *Adopter {
	return &Adopter{v220: v220}
}

// 业务类
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

// 普通的业务功能
func (p *Phone) Charge() {
	fmt.Println("Phone 进行了充电....")
	p.v.Use5V()
}

func main() {
	phone := NewPhone(NewAdopter(new(V220)))
	phone.Charge()
}
