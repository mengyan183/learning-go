package constant

import "testing"

/**
通过使用iota 实现连续常量赋值操作
iota: 常量计数器,只能在常量的表达式中使用;iota在每一个匹配的const出现时都会初始化为0;
在连续常量中,通常对第一个常量使用iota进行赋值操作,则后续的常量会根据 iota递增值 进行运算操作
*/
const (
	MONDAY    = iota + 1 // 0+1
	TUESDAY              //1+1
	WEDNESDAY            //2+1
	THURSDAY             //3+1
	FRIDAY               //4+1
	SATURDAY             //5+1
	SUNDAY               //6+1
)
const (
	Readable   = 1 << iota // 执行左移位运算  1 << 0  1(二进制)
	Writeable              // 1 << 1 10(二进制)
	Executable             // 1 << 2  100(二进制)
)

func TestConstant(t *testing.T) {
	t.Log(MONDAY, TUESDAY)
}

func TestConstantBitwise(t *testing.T) {
	t.Log(Readable, Writeable, Executable)
	i := 3 // 11(二进制)
	t.Log(Readable&i == Readable, Writeable&i == Writeable, Executable&i == Executable)
}
