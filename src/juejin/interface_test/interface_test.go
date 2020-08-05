package interface_test

import (
	"fmt"
	"testing"
)

type object interface {
	A()
	B()
}

type aImpl struct {
}

func (a aImpl) A() {
	fmt.Printf("%T\n", a)
}

func (a aImpl) B() {
	fmt.Printf("%T\n", a)
}

type bImpl struct {
}

func (a bImpl) A() {
	fmt.Printf("%T\n", a)
}

func (a bImpl) B() {
	fmt.Printf("%T\n", a)
}

func TestImpl(t *testing.T) {
	// 通过接口实现多态
	var o object
	o = new(aImpl)
	interfaceA(o)
	o = new(bImpl)
	interfaceA(o)
}

func interfaceA(o object) {
	o.A()
}
