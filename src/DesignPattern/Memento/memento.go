package main

import (
	"fmt"
)

type User struct {
	level   int
	attack  int
	defense int
	life    int
}

func (u *User) InitState() {
	u.attack = 1
	u.level = 1
	u.defense = 1
	u.life = 100
}

func (u *User) SuperState() {
	u.attack = 100
	u.level = 1
	u.defense = 100
	u.life = 1000
}

func (u *User) ShowState() {
	fmt.Println("attack:", u.attack)
	fmt.Println("level:", u.level)
	fmt.Println("defense:", u.defense)
	fmt.Println("life:", u.life)
}

func (u *User) SaveState() RoleStateMemento {
	return RoleStateMemento{u.level, u.attack, u.defense, u.life}
}

func (u *User) RecoverState(rsm RoleStateMemento) {
	u.attack = rsm.attack
	u.level = rsm.level
	u.defense = rsm.defense
	u.life = rsm.life
}

type RoleStateMemento struct {
	level   int
	attack  int
	defense int
	life    int
}

type RoleStateTaker struct {
	rsm RoleStateMemento
}

func main() {
	zhangchen := new(User)
	zhangchen.InitState()
	zhangchen.ShowState()

	zhangchen.SuperState()
	zhangchen.ShowState()

	rst := &RoleStateTaker{rsm: zhangchen.SaveState()}

	zhangchen.InitState()
	zhangchen.ShowState()

	zhangchen.RecoverState(rst.rsm)
	zhangchen.ShowState()
}
