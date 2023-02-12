package main

import "fmt"

// 观察者模式
// 假设现在有这样的场景
// 一群人在抄作业，班长在放哨，这时候老师来了
// 首先 学生应该有一个状态，并且有一个老师来了怎么办的方法
// 班长拥有所有学生的信息，并且在老师来了之后能通知所有的同学

// 抽象层

type Listener interface {
	OnTeacherComming() //观察者得到通知后要触发具体的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// 实现层
// 观察者具体的学生
type StuZhang3 struct {
	Badthing string
}

func (s *StuZhang3) OnTeacherComming() {
	fmt.Println("张三 停止...", s.Badthing)
}
func (s *StuZhang3) DoBadthing() {
	fmt.Println("张三 正在...", s.Badthing)
}

type StuZhaosi struct {
	Badthing string
}

func (s *StuZhaosi) DoBadthing() {
	fmt.Println("赵四 正在...", s.Badthing)
}
func (s *StuZhaosi) OnTeacherComming() {
	fmt.Println("赵四 停止...", s.Badthing)
}

type StuWangwu struct {
	Badthing string
}

func (s *StuWangwu) DoBadthing() {
	fmt.Println("王五 正在....", s.Badthing)
}
func (s *StuWangwu) OnTeacherComming() {
	fmt.Println("王五 停止...", s.Badthing)
}

// 班长

type ClassMonitor struct {
	listenerList []Listener // 需要通知的全部观察者集合
}

func (c *ClassMonitor) AddListener(listener Listener) {
	c.listenerList = append(c.listenerList, listener)
}

func (c *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range c.listenerList {
		// 找到要删除的位置
		if listener == l {
			c.listenerList = append(c.listenerList[:index], c.listenerList[index+1:]...)
			break
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, listener := range c.listenerList {
		listener.OnTeacherComming() // 多态现象
	}
}
func main() {
	s1 := &StuZhang3{
		Badthing: "抄作业",
	}
	s2 := &StuZhaosi{
		Badthing: "玩王者荣耀",
	}
	s3 := &StuWangwu{
		Badthing: "看赵四玩王者荣耀",
	}

	classMonitor := new(ClassMonitor)
	classMonitor.AddListener(s1)
	classMonitor.AddListener(s2)
	classMonitor.AddListener(s3)

	fmt.Println("上课了但是老师没来...")
	s1.DoBadthing()
	s2.DoBadthing()
	s3.DoBadthing()

	fmt.Println("现在老师来了...班长给同学们使了个眼神...")
	classMonitor.Notify()

}
