package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	/**
		变量声明方式
			var i int = 0
	============================
			i := 0
	============================
			var i int
			i = 0
	*/
	var i int = 1
	var j int = 1
	var total int = i + j
	t.Log(i)
	for q := 0; q < 5; q++ {
		t.Log(j)
		temp := j
		j = i + j
		i = temp
		total += j
	}
	t.Log(total)
}

func TestExchange(t *testing.T) {
	// 直接使用类型推断实现赋值操作
	a := 1
	b := 2
	/*temp := a
	a = b
	b = temp*/
	// 可以在一个赋值语句中实现多个变量赋值操作,因此可以实现变量交换操作等价于引用中间变量实现值交换操作
	a, b = b, a
	t.Log(a, b)
}
