package main

import "fmt"

// 代理模式
// 假如有一个抽象的购物主题
// 有韩国购物，美国购物，日本购物
// 我们设置一个代理

type Goods struct {
	Kind string //商品的种类
	Fact bool   //商品的真伪
}

// -------抽象层-----
type Shopping interface {
	Buy(goods *Goods) //某个任务
}

// ----实现层-----
type KoreaShopping struct{}

func (k *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("在韩国进行了购物.....")
}

type AmericanShopping struct{}

func (k *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("在美国进行了购物.....")
}

type AfricaShopping struct{}

func (k *AfricaShopping) Buy(goods *Goods) {
	fmt.Println("在非洲进行了购物.....")
}

// 海外的代理
type OverSeasProxy struct {
	shopping Shopping //代理某个服务
}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{shopping: shopping}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	// 1.先辨别真伪
	if op.distinguish(goods) {
		// 2.调用具体要被代理的购物方式的Buy()方法
		// 3.海关安检
		op.shopping.Buy(goods)
		op.check(goods)
	}

}

// 来个辨别真伪的能力
func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对{", goods.Kind, "}进行了辨别真伪")
	if !goods.Fact {
		fmt.Println("发现假货", goods.Kind, ",不要购买.")
	}
	return goods.Fact
}

// 来个检查的流程
func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对:", goods.Kind, "进行了安检")
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	kshopping := new(KoreaShopping)

	proxy := NewProxy(kshopping)
	proxy.Buy(&g1)
	proxy.Buy(&g2)
}
