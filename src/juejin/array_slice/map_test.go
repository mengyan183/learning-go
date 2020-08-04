package array_slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"unsafe"
)

func TestDeclareInit(t *testing.T) {
	// 只声明
	var m1 map[int]int // 初始化值为nil,但支持访问
	t.Log(m1)
	t.Log(m1[1]) // 输出了0
	t.Logf("%T , %p,%p", m1, m1, &m1)
	t.Log(unsafe.Pointer(&m1))
	// 判断key是否存在
	i, ok := m1[1]
	t.Log(i, ok)
	//m1[0] = 0
	// 删除map中的数据
	delete(m1, 0)
	// 获取当前map中数据个数
	l := len(m1)
	t.Log(l)
	// cap不支持nil map
	//c := cap(m1)

	for k, v := range m1 {
		t.Log(k, v)
	}

}

func TestZeroValue(t *testing.T) {
	type T struct {
		i    int
		f    float64
		next *T
	}
	var t1 T
	t.Log(t1)

}

func TestMapWithSlice(t *testing.T) {
	ms := make([]map[string]string, 0, 2)
	ms = append(ms, map[string]string{"m1": "v1"})
	ms = append(ms, map[string]string{"m2": "v2"})
	for _, v := range ms {
		for k, v := range v {
			t.Log(k, v)
		}
	}
}

func TestSyncMap(t *testing.T) {
	var synmap sync.Map
	// 数据写入
	synmap.Store("m1", "v1")
	synmap.Store("m2", "v2")

	// 数据读取
	v, ok := synmap.Load("m1")
	if ok {
		t.Log(v)
	}

	// 数据删除
	synmap.Delete("m1")
	// 数据遍历
	synmap.Range(func(key, value interface{}) bool {
		t.Log(key, value)
		return true
	})
}

func TestDeclareInitWithSource(t *testing.T) {
	//var m1 map[int]int
	//m1[1] = 1
	var m map[string]string
	//m["1"] = "1"
	//_,_=m["1"]
	_=m["1"]
}

func TestFn4(t *testing.T) {
	assert.Equal(t,5,f4())
}
func f4() (x int) {
	defer func(x int) {
		// 操作的是局部变量,而不是返回值x
		x++
		fmt.Println(x)
	}(x)// 这里压入栈的x的值为0
	return 5 // 这里的操作实际可以解析为 返回值x = 5 defer函数 return
}