package array_slice

import "testing"

// 对于所有值类型数据 进行数据拷贝时都是copy 数据本身
func TestDeepCopy(t *testing.T) {
	s := []int{1, 2, 3, 4}
	s1 := make([]int, 0)
	for _, v := range s {
		s1 = append(s1, v)
	}
	s[0] = 0
	t.Log(s1)
	s2 := make([]int, 0)
	// copy不会对目标切片进行自动扩容
	copy(s2, s)
	t.Log(s2)
}
