package main

import (
	"fmt"
)

type AbstractPerson interface {
	show()
}

type ConcreatePerson struct {
	Name string
}

func (p *ConcreatePerson) show() {
	fmt.Println("Dress ", p.Name)
}

type Decorator struct {
	AbstractPerson
}

func (d *Decorator) decorate(comp AbstractPerson) {
	d.AbstractPerson = comp
}

func (d *Decorator) show() {
	if d.AbstractPerson != nil {
		d.AbstractPerson.show()
	}
}

type TShirts struct {
	Decorator
}

func (t *TShirts) show() {
	fmt.Println("TShirts")
	t.Decorator.show()
}

type Sneaker struct {
	Decorator
}

func (s *Sneaker) show() {
	fmt.Println("Sneaker")
	s.Decorator.show()
}

func main() {
	cp := &ConcreatePerson{"xiao_cai"}
	fmt.Println("decorate")
	ts := new(TShirts)
	sn := new(Sneaker)
	ts.decorate(sn)
	sn.decorate(cp)
	ts.show()
}
