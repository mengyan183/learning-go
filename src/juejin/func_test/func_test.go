package func_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 函数
func TestFuncParam(t *testing.T) {
	// 传递匿名函数
	i := funcParam(func(i int) int {
		return i
	})
	// 执行返回的函数
	t.Log(i())
}

// 函数作为入参和返回值
func funcParam(f func(i int) int) func() int {
	i := 1
	r := f(i) // 这里称为funcParam的回调函数
	return func() int {
		return r
	}
}

func TestDefer(t *testing.T) {
	testDefer()
}

func testDefer() {
	i := 0
	// 当执行到defer 这个语句时,会将要延时执行的函数压入defer专用函数栈中
	// 对于以下代码实际是将
	// func(i int) func() {
	//		i++
	//		// 第一步
	//		fmt.Println(i)
	//		return func() {
	//			// 最后一步
	//			fmt.Println("in")
	//			fmt.Println(i)
	//		}
	//	}(i)()
	// 返回的函数压入栈中,为了得到返回的函数,因此需要先将最外部的匿名函数执行
	defer func(i int) func() {
		i++
		// 第一步
		fmt.Println(i)
		return func() {
			// 最后一步
			fmt.Println("in")
			fmt.Println(i)
		}
	}(i)()
	// 当存在多个defer 时,按照栈 先进后出的原则,后入栈的函数会先执行
	defer func(i int) {
		fmt.Println("后续i的变化不会影响已被压入栈的函数")
		fmt.Println(i)
	}(i)
	i++
	// 第二步
	fmt.Println("out")
	fmt.Println(i)

}

func deferReturn() (result int) {
	i := 0
	i++
	defer func() {
		// 这里相当于是对返回值再一次操作
		result++
	}()
	return i
}

func TestDeferReturn(t *testing.T) {
	i := deferReturn()
	assert.Equal(t, 2, i)
}

func getFn() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// 闭包测试
func TestGetFn(t *testing.T) {
	// 在返回的fn函数中实际 getFn()中的局部变量x并未消亡,而是随着返回的函数一起返回作为返回函数的全局变量
	fn := getFn()
	// 因此没调用一次fn ,都会执行x++操作
	t.Log(fn())
	t.Log(fn())
	// 而当重新执行getFn时,实际是重新生成了新的闭包,因此x又重置为0
	fn = getFn()
	t.Log(fn())
}
