package array_slice

import (
	"testing"
	"unsafe"
)

// 切片元素删除
func TestRmData(t *testing.T) {
	s := []int{1, 2, 3, 4}
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])))
	s = s[1:]
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])))

	s = append(s[:0], s[1:]...)
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])))
	s = []int{1, 2, 3, 4}
	// 删除指定下标元素
	i := 1
	s = append(s[:i], s[i+1:]...)
	t.Log(s)
}
