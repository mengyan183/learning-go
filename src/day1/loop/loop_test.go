package loop

import "testing"

/**
for 循环
*/
func TestLoop(t *testing.T) {
	n := 0
	// 类似于 while(n<5) 
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestInfiniteLoop(t *testing.T) {
	n := 0
	for {
		t.Log(n)
		n++ // 只有后置自增和自减操作
	}
}
