package main

import "fmt"

type Tv struct{}

func (t *Tv) On()  { fmt.Println("打开电视机") }
func (t *Tv) Off() { fmt.Println("关闭电视..") }

type Inxiang struct{}

func (i *Inxiang) On()  { fmt.Println("打开音响..") }
func (i *Inxiang) Off() { fmt.Println("关闭音响..") }

type Light struct{}

func (l *Light) On()  { fmt.Println("打开灯光..") }
func (l *Light) Off() { fmt.Println("关闭灯光..") }

// 外观类
type AControl struct {
	tv    *Tv
	in    *Inxiang
	light *Light
}

func (a *AControl) KTV() {
	fmt.Println("现在是ktv模式...")
	a.tv.On()
	a.light.Off()
	a.in.On()
}

func main() {
	// 比方说现在ktv了
	ac := new(AControl)
	ac.KTV()
}
