package main

//外观对象代理控制子系统的某系操作，
//需要时可以屏蔽掉子系统中的一些操作，使其不对外暴露
import (
	"fmt"
)

type BuySell interface {
	sell()
	buy()
}

type ohterOP interface {
	zhuanrang()
}

type Stock1 struct {
}

func (s *Stock1) sell() {
	fmt.Println("股票1卖出")
}

func (s *Stock1) buy() {
	fmt.Println("股票1买进")
}

func (s *Stock1) zhuanrang() {
	fmt.Println("股票1zr")
}

type Stock2 struct {
}

func (s *Stock2) sell() {
	fmt.Println("股票2卖出")
}

func (s *Stock2) buy() {
	fmt.Println("股票2买进")
}

func (s *Stock2) zhuanrang() {
	fmt.Println("股票2zr")
}

type other interface {
}

type NationalDebt1 struct {
}

func (n *NationalDebt1) sell() {
	fmt.Println("国债1卖出")
}

func (n *NationalDebt1) buy() {
	fmt.Println("国债1买进")
}

type Fund struct {
	gu1 BuySell
	gu2 BuySell
	nd1 BuySell
}

func (f *Fund) buyFund() {
	f.gu1.buy()
	f.gu2.buy()
	f.nd1.buy()
}

func (f *Fund) sellFund() {
	f.gu1.sell()
	f.gu2.sell()
	f.nd1.sell()
}

func NewFund() *Fund {
	return &Fund{&Stock1{}, &Stock2{}, &NationalDebt1{}}
}

type Fundohter struct {
	gu1 ohterOP
	gu2 ohterOP
}

func (f *Fundohter) zhuanrangFund() {
	f.gu1.zhuanrang()
	f.gu2.zhuanrang()
}

func NewFundOt() *Fundohter {
	return &Fundohter{&Stock1{}, &Stock2{}}
}

func main() {
	jijin := NewFund()
	jijin.buyFund()
	jijin.sellFund()
	
	zhuanrangFun :=NewFundOt()
	zhuanrangFun.zhuanrangFund()
}
