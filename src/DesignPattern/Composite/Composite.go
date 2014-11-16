package main

import (
	"fmt"
)

//可对Componet进行进一步的抽象
//如将用两个不同的基类Componet1,Componet2继承Componet
//让leaf继承Componet1，Composite继承Componet2
//从而对leaf 和 Composite的一些独有行为进行定义
type Componet interface {
	add(Componet)
	remove(Componet)
	display()
}

type leaf struct {
	Componet
	name string
}

func (l *leaf) add(c Componet) {
	fmt.Println("can't add")
}

func (l *leaf) remove(c Componet) {
	fmt.Println("can't remove")
}

func (l *leaf) display() {
	fmt.Println(l.name)
}

type Composite struct {
	Componet
	name string
	list map[Componet]interface{}
}

func (com *Composite) add(c Componet) {
	com.list[c] = nil
}

func (com *Composite) remove(c Componet) {
	delete(com.list, c)
}

func (com *Composite) display() {
	fmt.Println(com.name)
	for key, _ := range com.list {
		key.display()
	}
}

func NewComposite(name string) *Composite {
	return &Composite{list: make(map[Componet]interface{}), name: name}
}

func main() {
	l1 := &leaf{name: "l1"}
	l2 := &leaf{name: "l2"}
	l3 := &leaf{name: "l3"}

	com1 := NewComposite("com1")
	com2 := NewComposite("com2")
	com3 := NewComposite("com3")

	com1.add(com2)
	com1.add(l1)
	com2.add(l2)
	com2.add(com3)
	com3.add(l3)

	com1.display()
}
