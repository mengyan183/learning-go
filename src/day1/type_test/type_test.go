package type_test

import (
	"math"
	"testing"
)

/**
测试隐式数据类型
*/
type MyInt int64 // 自定义类型别名
func TestImplicitTypeConversion(t *testing.T) {
	/*
		// 当前机器是64位,实际当前a是64位数据 (int64)
		var a int = 1
		var b int64
		// 当执行赋值操作时,已经提示编译错误(cannot use a (type int) as type int64 in assignment)
		b = a
	*/
	/*
		// int32 不会自动转换为高位 int64,提示编译错误
			var a int32 = 1
			var b int64
			b = a
	*/
	/*	// 当自定义别名类型时,即使别名和实际类型相同,也不能直接进行隐式转换
		var a int64 = 1
		var b MyInt
		b = a*/

	// 对于类型转换,只能通过显式类型转换
	var a int = 1
	var b MyInt
	b = MyInt(a)
	t.Log(b)

}

/**
使用math包中的默认值
*/
func TestDefault(t *testing.T) {
	a := math.MaxInt8
	t.Log(a)
}

/**
指针
*/
func TestPoint(t *testing.T) {
	a := 1
	// 获取a的指针地址
	aPtr := &a
	// 指针不支持直接运算,提示编译错误
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	// 打印 数据类型
	t.Logf("%T %T", a, aPtr)
}

/**
字符串初始化
*/
func TestString(t *testing.T) {
	var s string
	// string是基本数据类型不能赋值 nil
	//s = nil //cannot use nil as type string in assignment
	// 字符串的初始化默认值为空字符串
	t.Log(s == "", len(s)) // true,0
}
