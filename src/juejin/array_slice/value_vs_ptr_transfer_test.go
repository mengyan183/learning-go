package array_slice

import (
	"testing"
	"unsafe"
)

// 值传递vs指针传递

// 值传递包含 int/float/string/bool/array/struct(非内置struct)

func TestValueTransfer(t *testing.T) {
	i := 1
	t.Log(unsafe.Pointer(&i))
	// 由于是值传递所以变量j指针内存地址会发生变化
	j := i
	t.Log(unsafe.Pointer(&j))

	s := "1"
	t.Log(unsafe.Pointer(&s))
	s1 := s
	t.Log(unsafe.Pointer(&s1))

	arr := [5]int{1}
	t.Log(unsafe.Pointer(&arr))
	arr1 := arr
	t.Log(unsafe.Pointer(&arr1))

	type custom struct {
	}

	c := custom{}
	t.Log(unsafe.Pointer(&c))
	c1 := c
	t.Log(unsafe.Pointer(&c1))

	type notEmptyCustom struct {
		name string
	}
	nc := notEmptyCustom{}
	t.Log(unsafe.Pointer(&nc))
	nc1 := nc
	t.Log(unsafe.Pointer(&nc1))
}

// 对于slice/map/chan/pointer 以及字段为空的empty struct 等都为地址传递
// 地址传递是体现在由于在这些容器对象内部保存的都是元素的内存地址,因此是引用传递
func TestPtrTransfer(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(unsafe.Pointer(&s))
	s1 := s
	t.Log(unsafe.Pointer(&s1))

	m := map[int]int{1: 1, 2: 2}
	t.Log(unsafe.Pointer(&m))
	m1 := m
	t.Log(unsafe.Pointer(&m1))

	c := make(chan int, 10)
	t.Log(unsafe.Pointer(&c))
	c1 := c
	t.Log(unsafe.Pointer(&c1))

	ptr := new(int)
	t.Log(unsafe.Pointer(&ptr))
	ptr1 := ptr
	t.Log(unsafe.Pointer(&ptr1))
}
