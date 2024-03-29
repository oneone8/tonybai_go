package samps

import (
	"fmt"
	"testing"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}
func Test2(t *testing.T) {

	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
