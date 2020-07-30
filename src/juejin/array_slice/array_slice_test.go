package array_slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

// 数组声明并自动初始化
func TestArrayDeclareInit(t *testing.T) {
	var declareArr [10]int
	t.Log(declareArr)
}

// 声明初始化并实例化
func TestArrayInstant(t *testing.T) {
	a := [10]int{1, 2, 3}
	t.Log(a, len(a))
	assert.Equal(t, len(a), cap(a))
	assert.Equal(t, 10, len(a))
	b := [10]int{0: 1, 9: 10}
	t.Log(b, len(b))
	c := [...]int{1, 2, 3}
	t.Log(c, len(c))
	assert.Equal(t, 3, cap(c))
	d := [...]int{10: 11}
	t.Log(d, len(d))
	assert.Equal(t, 11, cap(d))
}

func TestArrayEach(t *testing.T) {
	d := [...]int{10: 11}
	var dst2 []*int
	for _, v := range d {
		t.Log(v)
		dst2 = append(dst2, &v)
	}
	t.Log(dst2)
	var v int
	// 等价于 ; 表明对于range中的v实际是共享变量,在整个遍历过程中变量v的内存地址是一直不变的
	for i := 0; i < len(d); i++ {
		v = d[i]
		dst2 = append(dst2, &v)
	}
	t.Log(dst2)
}

func TestSliceDeclare(t *testing.T) {
	var s1 []int
	t.Log(s1, len(s1), cap(s1))
	// make 会 声明并自动初始化容器中的数据为当前数据类型的默认值
	s2 := make([]int, 1)
	t.Log(s2, len(s2), cap(s2))
	assert.Equal(t, len(s2), cap(s2))
	//编译报错 len larger than cap in make([]int)
	//s3 := make([]int, 2, 1)
	//t.Log(s3, len(s3), cap(s3))
	//assert.Equal(t, len(s3), cap(s3))
	s4 := make([]int, 2, 3)
	t.Log(s4, len(s4), cap(s4))
}

func TestSliceAppend(t *testing.T) {
	var s1 []int
	s1 = append(s1, 1, 2)
	t.Log(s1, len(s1), cap(s1))
	s2 := make([]int, 1)
	s2 = append(s2, s1...)
	t.Log(s2, len(s2), cap(s2))
}

func TestCompareNewVarMake(t *testing.T) {
	var s1 []int
	t.Log(unsafe.Pointer(&s1), s1 == nil)
	s2 := new([]int) // s2为指针
	t.Log(unsafe.Pointer(s2), *s2 == nil)
	s3 := make([]int, 1)
	t.Log(unsafe.Pointer(&s3), s3 == nil)

	_ = make(chan int, 5)
	m := make(map[int]string, 10)
	for k, v := range m {
		t.Log(k, ":", v, ";")
	}
	v, ok := m[0]
	t.Log(v, ok)
}
// 对于slice 而言,slice结构体的组成包含
//type slice struct {
//	array unsafe.Pointer
//	len   int
//	cap   int
//}
// 对于unsafe.Pointer获取到的是当前slice结构实例的内存地址,而%p 获取到的是当前slice中保存的指向其数组的指针
// 因此当slice扩容实际是对底层关联的数组进行扩容(复制迁移)操作,所以当扩容后slice的地址不会变化,变化的是slice关联的数组的地址
// 切片每次扩容后的容量都是之前容量的2倍
func TestSliceAutoIncr(t *testing.T) {
	// 当前slice的容量为1
	s := make([]int, 1)
	t.Log(unsafe.Pointer(&s),unsafe.Pointer(&(s[0])),len(s),cap(s))
	t.Logf("%p\n",s)
	s = append(s, 1)
	t.Log(unsafe.Pointer(&s),len(s),cap(s))
	t.Logf("%p\n",s)
	// 当扩容后,指针指向的数组内存地址会发生变化
	s = append(s, 2)
	t.Log(unsafe.Pointer(&s),len(s),cap(s))
	// 这里是调用 /usr/local/Cellar/go/1.14.5/libexec/src/reflect/value.go:1429
	t.Logf("%p\n",s)
}
