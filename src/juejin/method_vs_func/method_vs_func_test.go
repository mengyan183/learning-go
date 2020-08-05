package method_vs_func

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyInt int

func (i MyInt) customMethod(a int) int {
	return int(i) + a
}
func (iptr *MyInt) customMethodPtr(a int) int {
	return int(*iptr) + a
}

func TestMethod(t *testing.T) {
	i := MyInt(0)
	a := i.customMethod(1)
	assert.Equal(t, 1, a)
	a = i.customMethodPtr(2)
	assert.Equal(t, 2, a)
}

type Person struct {
	Age int
}

func (p Person) p() {
	fmt.Printf("%T\n", p)
}

type Animal struct {
}

func (a Animal) a() {
	fmt.Printf("%T\n", a)
}

type Man struct {
	// 嵌套匿名结构体; 通过匿名结构体实现继承的效果
	Person
	Animal
}

func (m Man)a()  {
	fmt.Printf("%T\n", m)
}

func TestStructExtend(t *testing.T) {
	m := new(Man)
	m.a() // 当重写了a()后实际调用的是重写后的a()
	m.p()
}
