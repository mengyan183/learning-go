package oop

import (
	"fmt"
	"testing"
	"time"
)

// 自定义类型
/// 自定义类型为方法
type IntConversionInt func(int) int

// 使用内置函数实现自定义
func SpendTime(f IntConversionInt) IntConversionInt {
	return func(i int) int {
		start := time.Now()
		ret := f(i)
		fmt.Printf("花费时间 : %d ms", time.Since(start).Milliseconds())
		return ret
	}
}

func ConversionInt(i int) int {
	time.Sleep(time.Second * 1)
	return i
}

func TestCustomFuncType(t *testing.T) {
	t.Log(SpendTime(ConversionInt)(1))
}
