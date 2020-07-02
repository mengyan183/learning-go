package condition

import "testing"

func TestCondition(t *testing.T) {
	// 支持在if 条件中使用变量定义与赋值操作; 主要原因在于函数支持多返回值
	if a := 1 == 1; a {
		t.Log("1==1")
		t.Log(a)
	}
	//// 由于函数支持多返回值,通过返回值来执行不同的操作
	//if a, e := someFunc(); e == nil {
	//	t.Log(a)
	//} else {
	//	t.Log(e)
	//}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2: // 可以设置多个匹配条件,通过 ","分割来实现多条件命中
			t.Log("Even") // 当命中case后,即使不显式声明break,也不会导致case击穿
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("It is not between 0-3")
		}
	}
}

func TestSwitchWithoutCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0: // 可以通过表达式的方式作为条件判断,可以达到 if condition elseif condition模式
			t.Log("Even") // 当命中case后,即使不显式声明break,也不会导致case击穿
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unKnown")
		}
	}
}
