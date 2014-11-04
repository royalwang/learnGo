package main

import ()

type AbstractPerson interface {
	show(string)
}

type Show func(string)

func (s Show) show(st string) {
	s(st)
}

func show1(s1 string){
    
}

func main() {
}
