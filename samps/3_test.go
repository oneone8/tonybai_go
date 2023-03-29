package samps

import "testing"

type TT struct {
	a int
}

func (t TT) M1() {
	t.a = 10
}

func (t *TT) M2() {
	t.a = 11
}

func Test3(tt *testing.T) {
	var t TT
	println(t.a) // 0

	t.M1()
	println(t.a) // 0

	p := &t
	p.M1()
	println(t.a)

	p.M2()
	println(t.a) // 11
}
