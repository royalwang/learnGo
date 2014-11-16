package main

import (
	"fmt"
)

type Componet interface {
	display()
}

type leaf struct {
	Componet
	name string
}

type ConcreateLeaf struct {
	leaf
}

func (l *ConcreateLeaf) display() {
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

type ConcreateComposite struct {
	Composite
}

func (com *ConcreateComposite) display() {
	fmt.Println(com.name)
	for key, _ := range com.list {
		key.display()
	}
}

func NewConcreateComposite(name string) *ConcreateComposite {
	return &ConcreateComposite{Composite{list: make(map[Componet]interface{}), name: name}}
}

func main() {
	l1 := &ConcreateLeaf{leaf{name: "l1"}}
	l2 := &ConcreateLeaf{leaf{name: "l2"}}
	l3 := &ConcreateLeaf{leaf{name: "l3"}}

	com1 := NewConcreateComposite("com1")
	com2 := NewConcreateComposite("com2")
	com3 := NewConcreateComposite("com3")

	com1.add(com2)
	com1.add(l1)
	com2.add(l2)
	com2.add(com3)
	com3.add(l3)

	com1.display()
}
