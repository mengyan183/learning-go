package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func square(i int) int {
	return i * i
}

func TestUnit(t *testing.T) {
	inputs := [...]int{1, 2, 3, 4}
	expects := [...]int{1, 4, 9, 16}
	for i := 0; i < len(inputs); i++ {
		//if expects[i] != square(inputs[i]) {
		//	t.Error("预期不一致")
		//	// Fatal或FailNow会直接终止当前程序运行
		//	//t.Fatal("预期不一致")
		//	//t.FailNow()
		//}

		// 采用断言
		assert.Equal(t,expects[i],square(inputs[i]),"预期不一致")
	}
}
