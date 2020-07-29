package condition

import "testing"

func TestSwitch(t *testing.T) {
	n := 10
	switch n {
	case 10:
		t.Log(10)
		// 只会击穿后续紧跟着的唯一一个case
		fallthrough
	case 11:
		t.Log("数据击穿", 11)
	case 12:
		t.Error("不该击穿")
		t.Log("12")
	}
}
