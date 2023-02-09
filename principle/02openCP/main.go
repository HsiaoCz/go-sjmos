package main

import "fmt"

// 没有开闭原则
// 假设有一个银行职员的类
// 具有转账，贷款等方法
// type Banker struct{}

// 存款业务
// func (b *Banker) Save() {
// 	fmt.Println("进行了存款业务...")
// }

// 转账业务
// func (b *Banker) Transfor() {
// 	fmt.Println("进行了转账业务...")
// }

// 进行了支付业务
// func (b *Banker) Pay() {
// 	fmt.Println("进行了存款业务...")
// }

// 如果现在想添加一个新功能
// 比如股票的功能
//
//	func (b *Banker) Stack() {
//		fmt.Println("进行了股票业务...")
//	}
//
// 平铺设计的问题在于改动类本身
// 很难确定这种改动会不会对类的其他方法产生影响
func main() {
	// banker := &Banker{}
	// banker.Save()
	// banker.Transfor()
	// banker.Pay()

	// 存款的业务
	BankBusiness(&SaveBanker{})
}

// 我们可以抽象出来一个Banker 这个Banker有一个业务处理能力
// 每一个业务，单独创建一个Banker 继承这个能力，并对其重写

type Banker interface {
	Do()
}

// 实现一个抽象的架构层进行业务封装

func BankBusiness(banker Banker) {
	//通过接口向下调用
	banker.Do()
}

// 存款的banker
type SaveBanker struct{}

func (s *SaveBanker) Do() {
	fmt.Println("进行了存款....")
}

// 转账的banker
type TransBanker struct{}

func (t *TransBanker) Do() {
	fmt.Println("进行了转账....")
}

// 支付的banker
type PayBanker struct{}

func (p *PayBanker) Do() {
	fmt.Println("进行了支付....")
}

//这时候新增一个股票的业务
//新增一个股票的banker

type StackBanker struct{}

func (c *StackBanker) Do() {
	fmt.Println("进行了转账....")
}
