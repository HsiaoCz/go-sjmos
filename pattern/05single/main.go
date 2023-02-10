package main

import "fmt"

/*
  单例模式的三个要点:
        一、某个类只能有一个实例；
		二、它必须自行创建这个实例；
		三、它必须自行向整个系统提供这个实例
*/

// 保证一个类永远只能有一个对象，这个对象还能被系统的其他模块使用

// 1.保证这个类非公有化，外界不能通过这个类直接创建一个对象
// 这个类 类名应该小写
type singleton struct{}

// 2.但是保证还要有一个指针可以指向这个唯一对象，但是这个指针永远也不能改变方向
// 将指针私有化 不能通过这个指针私有化不让外界访问
var instance = new(singleton)

// 3.如果全部私有，那么外界模块永远访问这个对象
// 所以要向外提供一个方法来获取这个对象
// GetInstance 是否可以定义为singelton一个成员方法
// 答案是不能的 如果为成员方法 就必须要先访问对象，再访问函数
// 但是类和对象目前都已经私有化，外界无法访问，所以这个方法一定是一个全局普通函数

func GetInstance() *singleton {
	return instance
}

// 单例的某个方法
func (s *singleton) SomeThing() {
	fmt.Println("hello")
}

func main() {

	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()
	s2.SomeThing()
}

// 不过这种是饿汉式的单例
// 因为main函数还没执行之前 就已经分配内存了