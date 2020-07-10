package oop

import (
	"fmt"
	"testing"
)

// 声明一个接口
type Programmer interface {
	// 定义接口中的方法约束
	Code()
}

type GoProgrammer struct {
}

// 实现接口中的方法
// duckType式 实现接口方法,由于行为的方法约束和接口中的方法约束一致,因此认为当前结构体实现了接口的方法
func (p *GoProgrammer) Code() {
	fmt.Println("执行Code方法")
}

func (p *GoProgrammer) Write() {
	fmt.Println("执行Write方法")
}

func TestCode(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	p.Code()
}

// 接口变量
func TestInterfaceVariable(t *testing.T) {
	// 声明一个接口变量
	// 在当前接口变量p中实际存在两部分空间  1:类型 2:数据
	// p中类型对应的是 struct GoProgrammer
	// p中数据对应的是 &GoProgrammer{}
	var p Programmer = &GoProgrammer{}
	p.Code()
	// 编译错误 : cannot use GoProgrammer literal (type GoProgrammer) as type Programmer in assignment:
	//	GoProgrammer does not implement Programmer (Code method has pointer receiver)
	// 对于接口变量,由于GoProgrammer只是实现了Code方法,并没有强依赖关系,因此接口变量不能指向结构体的实例,而只能使用指向结构体实例化的指针
	//var p1 Programmer = GoProgrammer{}
}
