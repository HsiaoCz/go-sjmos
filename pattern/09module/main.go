package main

import "fmt"

// 抽象类，制作饮料，包裹一个模板、全部实现的步骤
type Beverage interface {
	// 煮开水
	BoilWater()

	// 冲突
	Brew()

	// 倒入茶叶
	Pour()

	// 添加佐料
	AddThings()
}

// 制作一套流程模板基类，让具体的制作流程继承且实现
type template struct {
	b Beverage
}

// 封装的固定模板，制作饮料
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.Pour()
	t.b.AddThings()
}

// 具体的制作流程
type MakeCoffee struct {
	template // 继承模板
}

// 煮水
func (mc *MakeCoffee) BoilWater() { fmt.Println("将水煮到100摄氏度....") }

// 冲泡
func (mc *MakeCoffee) Brew() { fmt.Println("用水冲咖啡豆...") }

// 导入杯中
func (mc *MakeCoffee) Pour() { fmt.Println("将冲好的咖啡导入杯中...") }

// 添加一些小料
func (mc *MakeCoffee) AddThings() { fmt.Println("加点小料...") }

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)

	// b 是Beverage 是MakeCoffee的接口，这里需要接口赋值，让b指向具体的子类
	// 来触发b的全部方法的多态特性
	makeCoffee.b = makeCoffee

	return makeCoffee
}

func main() {
	// 制作一杯咖啡
	makeCoffee := NewMakeCoffee()

	// 调用模板方法
	makeCoffee.MakeBeverage()
}
