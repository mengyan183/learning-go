package operator

import "testing"

/**
运算符
*/

/**
数组对比操作
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 4}
	c := [...]int{1, 2, 4, 3}
	d := [...]int{1, 2, 3, 4}

	//t.Log(a == b) // 编译错误,由于 a 和 b的数组长度不同 invalid operation: a == b (mismatched types [4]int and [3]int)
	t.Log(a == c) // false , 由于a 和 c 对应索引的元素值不相同,所以 a != b
	t.Log(a == d) // true
}

const (
	Readable   = 1 << iota // 执行左移位运算  1 << 0  1(二进制)
	Writeable              // 1 << 1 10(二进制)
	Executable             // 1 << 2  100(二进制)
)

/**
按位清零
将 &^ 左边数据相异位保留,相同位清零
*/
func TestBitClear(t *testing.T) {
	a := 7            //111
	a = a &^ Readable // 111 &^ 001 = 110
	t.Log(a)
	a = a &^ Writeable // 110 &^ 010 = 100
	t.Log(a)
	a = a &^ Executable // 100 &^ 100 = 0
	t.Log(a)

	i := 7
	j := 8
	t.Log((i &^ j) == 0)
	q := 7
	t.Log((i &^ q) == 0)
}
