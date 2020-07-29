package condition

import (
	"testing"
)

func TestForBreakCustomPoint(t *testing.T) {
flg:
	for i := 0; i < 10; i++ {
	flg2:
		for j := 0; j < 10; j++ {
			for p := 0; p < 10; p++ {
				//if j%2 == 1 {
				if p%2 == 1 { // 导致 if i%2 == 1 永远不可达
					// 终止循环,并跳转到指定位置
					t.Log("执行flg2终止")
					break flg2
				}
			}
			if i%2 == 1 {
				// 终止循环,并跳转到指定位置,执行标签代码块后的代码
				t.Log("执行flg终止")
				break flg
			}
		}
		t.Log(i)
	}

}

func TestGotoLabel(t *testing.T) {
label:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for p := 0; p < 10; p++ {
				if p%2 == 1 {
					// 终止循环,并跳转到指定位置
					t.Log("执行flg2终止")
					goto label // 不再执行后续代码,跳转到指定标签,并继续执行标签的代码块
				}
			}
		}
		t.Log(i)
	}
}

func TestGotoLabel2(t *testing.T) {
TestLabel: //标签
	for a := 20; a < 35; a++ {
		if a == 25 {
			a += 1
			goto TestLabel
		}
		t.Log(a)
		a++
	}
}
