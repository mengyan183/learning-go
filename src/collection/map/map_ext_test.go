package _map

import (
	"testing"
)

// 设置 map的value为 func
func TestFuncVal(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}

	t.Log(m[1](2), m[2](2), m[3](2))
}

// 自定义实现set
func TestCustomSet(t *testing.T) {
	// 自定义实现set要求map的value为布尔类型
	customSet := map[int]bool{}
	// 对于添加数据,要求value必须都为true,否则无法直接判断元素是否存在于map中
	customSet[1]=true
	verifyExist(1,customSet,t)
	// 删除map中的元素
	delete(customSet,1)
	verifyExist(1,customSet,t)
	t.Log(len(customSet))
}

func verifyExist(n int, customSet map[int]bool,t *testing.T)  {
	if customSet[n]{
		t.Logf("%d 存在",n)
	}else {
		t.Logf("%d 不存在",n)
	}
}
