package main
/*
这种写法比较适合于输入参数统一的一批策略（函數）
*/
import (
	"fmt"
)

type input struct {
	money  float64
	Rebate float64
}

type cashSuper interface {
	AcceptCash(*input) float64
}

type cashHandler func(*input) float64

func (handler cashHandler) AcceptCash(i *input) (o float64) {
	o = handler(i)
	return
}

func cashNormal(i *input) (o float64) {
	return i.money
}

func cashRebate(i *input) (o float64) {
	return i.money * i.Rebate
}

type CashContext struct {
	cashSuper
}

func NewCashContext(str string) *CashContext {
	cash := new(CashContext)
	switch str {
	case "正常收费":
		cash.cashSuper = cashHandler(cashNormal)
	case "打8折":
		cash.cashSuper = cashHandler(cashRebate)
	}
	return cash
}

func main() {
	var total float64
	context := NewCashContext("正常收费")
	total += context.AcceptCash(&input{money: 10000})
	context = NewCashContext("打8折")
	total += context.AcceptCash(&input{money: 10000, Rebate: 0.8})
	fmt.Println(total)
}
