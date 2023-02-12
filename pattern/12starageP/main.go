package main

import "fmt"

// 武器策略
type WeaponStrategy interface {
	UseWeapon() // 使用武器
}

// 具体的策略
type Ak47 struct{}

func (ak *Ak47) UseWeapon() { fmt.Println("使用AK47 去战斗") }

// 具体的策略
type Knife struct{}

func (k *Knife) UseWeapon() { fmt.Println("使用小刀去战斗...") }

// 环境类 使用武器策略的类
type Hero struct {
	strategy WeaponStrategy //拥有一个抽象的策略
}

// 设置一个策略的方法
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	hero := Hero{}

	// 使用策略
    hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	// 使用策略
	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
