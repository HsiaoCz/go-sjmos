package main

import "fmt"

// 不用依赖倒转来实现需求
// 司机 张三  李四  汽车 奔驰 宝马
// 张三开奔驰 李四 开宝马

//type Benz struct {
//}
//
//func (b *Benz) Run() {
//	fmt.Println("Benz is running...")
//}
//
//type BMW struct {
//}
//
//func (b *BMW) Run() {
//	fmt.Println("BMW is running...")
//}

// 司机bob

//type Bob struct {
//}
//
//func (b *Bob) Drive(benz *Benz) {
//	fmt.Println()
//	benz.Run()
//}

// 司机alex

//type Alex struct {
//}
//
//func (a *Alex) Drive(bmw *BMW) {
//	fmt.Println()
//	bmw.Run()
//}

func main() {
	// bob开奔驰

	//benz := &Benz{}
	//bob := &Bob{}
	//bob.Drive(benz)

	// alex开宝马
	//bmw := &BMW{}
	//alex := &Alex{}
	//alex.Drive(bmw)

	// 如果bob开宝马
	// 这时候需要添加一个新的方法 给张三加一个drive bmw的方法
	// alex需要添加一个drive benz的方法
	// 当每次新添加一个司机或者汽车时
	// 都需要跟别的模块相耦合

	// bob开奔驰
	// 与抽象层依赖
	// bob 与 benz 都与driver car依赖
	// 指向具体的Bob对象 和 Car对象
	var benz Car
	var bob Driver
	benz = new(Benz)
	bob = new(Bob)
	bob.Drive(benz)

	var alex Driver
	var bmw Car
	bmw = new(BMW)
	alex = new(Alex)
	alex.Drive(bmw)

	alex.Drive(benz)
}

// 使用依赖倒转原则

// 我们在设计一个系统的时候 可以分为三层
// 业务逻辑层 司机驾驶汽车
// 抽象层  抽象的汽车 抽象的司机
// 实现层 汽车实现启动的方法 司机实现开车的方法

// ---抽象层----

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// ----实现层---

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("benz is running...")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("bmw is running...")
}

type Bob struct {
}

func (b *Bob) Drive(car Car) {
	fmt.Println("bob is driving")
	car.Run()
}

type Alex struct {
}

func (a *Alex) Drive(car Car) {
	fmt.Println("alex is driving")
	car.Run()
}
