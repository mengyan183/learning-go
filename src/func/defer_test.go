package _func

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// 在方法执行结束之后肯定会执行; 类似于 java 中的 try{}catch{}finally{} 中的finally操作
	defer Clear()
	t.Log("正常输出")
	// panic代表上抛异常
	panic("异常输出") // defer仍然会执行
	// panic后的代码实际是不可达的
}

func Clear() {
	fmt.Print("执行Clear")
}
