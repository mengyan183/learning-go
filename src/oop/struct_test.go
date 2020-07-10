package oop

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构体定义
type Employee struct {
	Id   int
	Name string
	Age  int
}

func TestInstantiate(t *testing.T) {
	e := Employee{1, "e", 18}
	t.Log(e)
	// 获取类型
	t.Logf("%T", e)
	e1 := Employee{
		Name: "e1",
		Age:  18,
	}
	t.Log(e1)
	t.Logf("%T", e1)
	// 等同于 &e ,这里的e2是指向当前实例化Employee的指针
	e2 := new(Employee)
	e2.Age = 28
	e2.Id = 2
	e2.Name = "e2"
	t.Log(e2)
	t.Logf("%T", e2)
}

// 参数为实例的toString方法(行为)
func (e Employee) InstanceToString() {
	fmt.Println("形参为实例")
	fmt.Printf("Name的地址为: %x\n", unsafe.Pointer(&e.Name))
	fmt.Printf("传入参数的地址为: %x\n", unsafe.Pointer(&e))
	fmt.Printf("Id:%d,Name:%s,Age:%d\n", e.Id, e.Name, e.Age)
}

// 参数为指针的toString方法(行为)
func (e *Employee) PointerToString() {
	fmt.Println("形参为指针")
	fmt.Printf("Name的地址为: %x\n", unsafe.Pointer(&e.Name))
	fmt.Printf("传入参数的地址为: %x\n", unsafe.Pointer(e))
	fmt.Printf("Id:%d,Name:%s,Age:%d\n", e.Id, e.Name, e.Age)
}

func TestEmployeeToString(t *testing.T) {
	e := Employee{1, "e", 18}
	ePointer := &e
	fmt.Printf("Name:地址为: %x\n", unsafe.Pointer(&(ePointer.Name)))
	fmt.Printf("地址为: %x\n", unsafe.Pointer(ePointer))
	// 使用实例和指针都可以直接调用结构体的行为,不论行为的入参是实例或指针
	// 对于调用 参数为实例 的 结构体方法时,实际会对实例中字段进行值复制操作,并开辟新的内存空间
	e.InstanceToString()
	ePointer.InstanceToString()
	fmt.Println("=======================")
	// 对于调用 参数为指针 的 结构体方法时,实际是使用的指针指向的地址中的数据,而不不会执行值复制操作
	e.PointerToString()
	ePointer.PointerToString()
}
