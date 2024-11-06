package main

import (
	"fmt"
)

type tem struct {
	C float64
	F float64
}

func (t *tem) CtoF() {
	t.F = (t.C * 9 / 5) + 32
}
func (t *tem) FtoC() {
	t.C = (t.F - 32) * 5 / 9
}
func main() {
	a := tem{C: 10}
	a.CtoF()
	fmt.Println(a)
}
