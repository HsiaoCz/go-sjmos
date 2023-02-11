package main

import (
	"fmt"
)

// 适配的目标
type Attack interface {
	Fight()
}

// 具体的技能
type Dabaojian struct{}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了大宝剑技能,将敌人击杀...")
}

type Hero struct {
	Name   string
	attack Attack //攻击方式
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能...")
	h.attack.Fight() //使用具体的战斗方式
}

// 适配者
type PowerOff struct{}

func (p *PowerOff) ShutDown() { fmt.Println("电脑即将关机....") }

// 适配器 适配电脑即将关机
type Adopter struct {
	poweroff *PowerOff
}

func (a *Adopter) Fight() {
	a.poweroff.ShutDown()
}

func NewAdopter(p *PowerOff) *Adopter {
	return &Adopter{poweroff: p}
}

func main() {
	// gailun := Hero{Name: "盖伦", attack: new(Dabaojian)}
	// gailun.Skill()
	gailun := Hero{Name: "盖伦", attack: NewAdopter(new(PowerOff))}
	gailun.Skill()
}
