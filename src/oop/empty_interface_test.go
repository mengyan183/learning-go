package oop

import (
	"fmt"
	"testing"
)

// 空接口

func DoSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unKnown:", v)
	}
	// switch等价于下面的ifelse语句
	if v, ok := i.(int); ok {
		fmt.Println("int:", v)
	} else if v, ok := i.(string); ok {
		fmt.Println("string:", v)
	} else {
		fmt.Println("unKnown:", v)
	}
}

func TestEmptyInterfaceAssert(t *testing.T) {
	DoSomething(1)
	DoSomething("1")
	DoSomething(false)

}
