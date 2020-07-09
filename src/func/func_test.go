package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回多个值
func ReturnMutiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestReturnMutiVal(t *testing.T) {
	//r1, r2 := ReturnMutiValues()
	//t.Log(r1, r2)
	// 可以使用 "_" 接收返回值,表示不需要使用当前返回的值
	r1, _ := ReturnMutiValues()
	t.Log(r1)
}

// 函数作为参数传递和返回值 ; 形参可以不用写
//func SpentTime(op func(o int) int) func(o int) int {
func SpentTime(op func(int) int) func(int) int {
	// 类似于AOP 面向切面编程
	// 执行自定义内置函数
	return func(n int) int {
		// 获取开始时间
		start := time.Now()
		// 执行具体传入的函数
		ret := op(n)
		// 打印传入函数执行胡斐的时间
		fmt.Println("花费的时间为:", time.Since(start).Milliseconds(), "ms")
		// 返回真正执行的函数
		return ret
	}
}

func Sleep(op int) int {
	// 线程sleep 1s
	time.Sleep(time.Second * 1)
	return op
}

func TestFuncAop(t *testing.T) {
	// 返回具体要执行的函数; 函数作为变量的值
	op := SpentTime(Sleep)
	i := op(10)
	t.Log(i)
	t.Log(SpentTime(Sleep)(11))
}

func TestSliceAsParam(t *testing.T) {
	// 定义切片
	s := []int{1, 2, 3, 4, 5}
	t.Log(s, len(s), cap(s))
	t.Logf("传参前 : %p", s)
	SliceParam(s, t)
	t.Log(s, len(s), cap(s))
	t.Logf("传参后 : %p", s)
}

func SliceParam(s []int, t *testing.T) {
	if s == nil || len(s) == 0 {
		t.Log("切片不能为空")
		return
	}
	// 修改切片中最后一个索引位置的值
	s[len(s)-1] = 999
	t.Logf("形参地址 : %p", s)
	s = append(s, 1000)
	t.Log(s, len(s), cap(s))
	t.Logf("执行append后的地址 : %p", s)
}

// 可变长度参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarSum(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}
