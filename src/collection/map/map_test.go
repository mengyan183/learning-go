package _map

import "testing"

// map声明
func TestMapDeclare(t *testing.T) {
	// 声明并初始化一个空的map
	m := map[string]int{}
	t.Log(m, len(m))
	// 添加或更新数据操作
	m["1"] = 1
	t.Log(m, len(m))
	// 声明并初始化赋值
	m1 := map[string]int{"1": 1, "2": 2}
	t.Log(m1, len(m1), m1["2"])
	// 使用make声明一个新的map
	m2 := make(map[string]int, 10) // 这里的10代表的是map的容量
	t.Log(m2, len(m2))
}

func TestNotExistKey(t *testing.T) {
	m := map[int]int{}
	m[1] = 1
	// 获取一个map中不存在的key
	a := m[2]
	// 当访问一个不存在的key时,不会返回nil,而会返回value类型的默认值
	t.Log(a)
	// 如果判断map中是否存在一个key
	if v, flag := m[2]; flag {
		t.Log(v)
	} else {
		t.Log("当前key不存在于map中")
	}
	n := map[int]string{}
	t.Log(n[0])
}

// 遍历map
func TestRangeMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		t.Log(k, v)
	}
}
