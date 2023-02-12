package main

import "fmt"

// 假设有这样一个场景
// 去路边烤串
// 有烤羊肉，烤鸡翅命令，有烤串师傅，和服务员
// 根据命令模式，实际烤串场景

type Shifu struct{}

func (s *Shifu) KaoYang() {
	fmt.Println("师傅烤了羊肉...")
}

func (s *Shifu) KaoJichi() {
	fmt.Println("师傅烤了鸡翅...")
}

func (s *Shifu) KaoQiezi() {
	fmt.Println("师傅烤了茄子...")
}

// 抽象的命令
type Cammand interface {
	Kao()
}

type CmdKaoYang struct {
	shifu *Shifu
}

func (cy *CmdKaoYang) Kao() {
	cy.shifu.KaoYang()
}

type CmdKaoJichi struct {
	shifu *Shifu
}

func (cy *CmdKaoJichi) Kao() {
	cy.shifu.KaoJichi()
}

type CmdKaoQiezi struct {
	shifu *Shifu
}

func (cy *CmdKaoQiezi) Kao() {
	cy.shifu.KaoQiezi()
}

// 服务员
type Fuwuyuan struct {
	cmdList []Cammand
}

func (f *Fuwuyuan) Notify() {
	if f.cmdList == nil {
		return
	}
	for _, cmd := range f.cmdList {
		cmd.Kao()
	}
}

func main() {
	// 具体的烧烤业务
	shifu := new(Shifu)
	// 烧烤命令
	kaoy := CmdKaoYang{shifu: shifu}
	kaoj := CmdKaoJichi{shifu: shifu}
	kaoq := CmdKaoQiezi{shifu: shifu}

	fw := new(Fuwuyuan)
	fw.cmdList = append(fw.cmdList, &kaoj, &kaoy, &kaoq)
	fw.Notify()
}
