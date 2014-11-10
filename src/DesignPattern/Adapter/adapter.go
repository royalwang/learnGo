package main

import (
	"fmt"
)

type TriSocket interface {//三角插座
	Charge1()
}

type DouSocket interface {//两脚插座
	Charge2()
}

type Wall struct {//墙，有一个两脚插座，无法插MAC，用Panel做适配
	ds DouSocket
}

type Panel struct {//插头为两脚插座（实现了相应接口）
	ts TriSocket
	ds DouSocket
}

func (p *Panel) Charge2() {
	if p.ts != nil {
		p.ts.Charge1()
	}

	if p.ds != nil {
		p.ds.Charge2()
	}

}

type Mac struct {
}

func (m *Mac) Charge1() {
	fmt.Println("Mac 通电")
}

func main(){
  m :=new(Mac)
  w:=new(Wall)
  p :=new(Panel)
  w.ds = p
  p.ts = m
  
  w.ds.Charge2()
}
