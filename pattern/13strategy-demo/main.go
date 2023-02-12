package main

import "fmt"

// 现在有这样一个场景
// 商城打折
// 策略A (0.8折) 策略B(消费200,返现100)，用策略模式模拟场景

// 销售策略
type MarketStrategy interface {
	GetPrice(price float64) float64
}

type StrategyA struct{}

func (s *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A..所有商品打八折...")
	return price * 0.8
}

type StrategyB struct{}

func (s *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B...满200减少100")
	if price >= 200 {
		return price - 100
	}
	return price
}

// 环境类

type Goods struct {
	Price    float64
	Strategy MarketStrategy
}

func (g *Goods) SetStrategy(m MarketStrategy) {
	g.Strategy = m
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值...", g.Price, " . ")
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	nike := Goods{
		Price: 210,
	}

	// 上午，商城策略A
	fmt.Println("上午商城执行策略A....")
	nike.SetStrategy(new(StrategyA))
	fmt.Println("上午nike鞋卖", nike.SellPrice())

	// 下午商城执行策略B
	fmt.Println("上午商城执行策略B....")
	nike.SetStrategy(new(StrategyB))
	fmt.Println("下午niki鞋卖", nike.SellPrice())
}
