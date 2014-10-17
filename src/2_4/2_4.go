package main

import (
	"fmt"
)

func main() {
	k:=example1(1)
	fmt.Print(k)
	example2(1)
	example3(1)
}

func example1(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func example2(Num int) {
	switch {
	case 0 <= Num && Num <= 3:
		fmt.Printf("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Printf("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Printf("7-9")
	}
}

func example3(x int) {
	a := true
J:
	for j := 0; j < 5; j++ {
		if a {
			for i := 0; i < 10; i++ {
				if i > 5 {
					a = false
					break J
				}
				fmt.Println(i)
			}
		}

	}

}
