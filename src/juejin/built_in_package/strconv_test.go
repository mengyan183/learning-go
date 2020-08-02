package built_in_package

import (
	"strconv"
	"testing"
)

func TestStrConv(t *testing.T) {
	// 由于作为强类型语言,不同类型之间的数据不能直接操作
	// 字符串转bool类型
	f, err := strconv.ParseBool("true")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%T ,%t", f, f)
	}
	// 字符串转int; 第一个10代表了进制,第二个64代表了位数
	i, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%T,%d", i, i)
	}

	// 整型和字符串之间相互转换
	s := strconv.Itoa(10)
	t.Log(s)
	v, err := strconv.Atoi(s)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%T,%d", v, v)
	}
}
