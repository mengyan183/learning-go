package _func

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestDefer(t *testing.T) {
	// 在方法执行结束之后肯定会执行; 类似于 java 中的 try{}catch{}finally{} 中的finally操作
	defer Clear()
	t.Log("正常输出")
	// panic代表上抛异常
	panic("异常输出") // defer仍然会执行
	// panic后的代码实际是不可达的
}

func Clear() {
	fmt.Print("执行Clear")
}

func GetFunc() func() {
	fmt.Print("[outside]")
	return func() {
		fmt.Print("[inside]")
	}
}

// 输出结果为 [outside][here][inside]
func TestDeferGetFunc(t *testing.T) {
	defer GetFunc()()
	fmt.Print("[here]")
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestRange(t *testing.T) {
	c := make(chan int, 5)
	go fibonacci(cap(c), c)

	for i := range c {
		assert.Equal(t, i, 3)
		fmt.Println(i)
	}
}

type Person struct {
	Name string
}

func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func TestEmptyStructInstance(t *testing.T) {
	var p Person
	t.Log(p == Person{}, unsafe.Pointer(&p), unsafe.Pointer(&Person{}))
	p.test()
}

type Man struct {
	Age int
}

func TestStruct(t *testing.T) {
	p := Person{"xingguo"}
	//m := Man{10}
	//p, _ = m.(Person)
	t.Log(p)
}

func TestCap(t *testing.T) {
	//i := 1
	// 该操作编译报错
	//j := i++
	//t.Log(j)
}

// 相同作用域的变量不可重复定义

var a = 0

//var a = 1
func TestDefineVariable(t *testing.T) {
	a := 1
	//a := 1
	t.Log(a)
}

const acon = 1

const (
	n1 = iota
	n2
	n3
	n4 = "test"
	n5
)

func TestConstIota(t *testing.T) {
	t.Log(n1)
	t.Log(n5)
}

func TestChar(t *testing.T) {
	v1 := 'A'
	v2 := "A"
	fmt.Printf("v1的类型是%T,%d\n", v1, v1) // int32 65
	fmt.Printf("v1的类型是%T,%s\n", v2, v2) //string A
}



