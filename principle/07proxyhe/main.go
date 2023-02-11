package main

import "fmt"

// 代理模式和谐版

// 抽象的主题
type BeautyWoman interface {
	//对男人抛媚眼
	MakeEyesWithMan()
	//和男人共度良宵
	HappyWithMan()
}

// 具体的主题
// panjinlian
type PanJinlian struct{}

// 抛媚眼
func (p *PanJinlian) MakeEyesWithMan() { fmt.Println("给西门大官人抛了个媚眼") }

// 和男人共度美好的时刻
func (p *PanJinlian) HappyWithMan() { fmt.Println("和西门大官人共度了美好的时光...") }

// 中间的代理人
type WangPo struct {
	woman BeautyWoman
}

func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman: woman}
}

// 抛媚眼
func (w *WangPo) MakeEyesWithMan() { w.woman.MakeEyesWithMan() }

// 和男人共度美好的时刻
func (w *WangPo) HappyWithMan() { w.woman.HappyWithMan() }

// 业务逻辑 西门大官人

func main() {
	// 大官人想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinlian))
	// 王婆命令金莲
	wangpo.MakeEyesWithMan()
	// 王婆带着大官人
	wangpo.HappyWithMan()
}
