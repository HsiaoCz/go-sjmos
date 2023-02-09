package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭...")
}

// 使用继承的方式给小猫
// 继承 使用匿名嵌套结构体的方式 匿名结构体要放在最后一个

// CatB 给小猫添加一个睡觉的方法
//type CatB struct {
//	Cat
//}
//
//func (c *CatB) Sleep() {
//	fmt.Println("小猫在睡觉...")
//}

// 使用组合的方式给小猫添加一个睡觉的方法

type CatC struct {
	// 还可以只让方法依赖 而类本身不依赖
	//cat *Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫在睡觉...")
}
func (cc *CatC) Eat(c *Cat) {
	fmt.Println("小猫在睡觉...")
}

func main() {
	c := &Cat{}
	c.Eat()

	fmt.Println("----------")

	//cb := &CatB{}
	//cb.Sleep()
	cc := &CatC{}
	cc.Eat(c)
}

// 使用组合和继承的方式有本质的不同
// 继承的方式 类直接指向了被继承的类的方法
// 组合的方式 类指向的被继承的类 被继承的类指向了该方法
