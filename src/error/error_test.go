package error

import (
	"errors"
	"os"
	"testing"
)

// 预制的错误
var customError = errors.New("输入参数必须等于1")

// 通过多返回值,实现错误机制

//斐波那契数列
func FibonacciSequence(i int) ([]int, error) {
	if i < 2 {
		return nil, customError
	}
	// 初始化数列
	fibonacciList := []int{1, 1}
	for j := 2; j < i; j++ {
		fibonacciList = append(fibonacciList, fibonacciList[j-1]+fibonacciList[j-2])
	}
	return fibonacciList, nil
}

func TestError(t *testing.T) {
	for i := 0; i < 10; i++ {
		if v, err := FibonacciSequence(i); err == nil {
			t.Log(v)
		} else {
			// 通过预制错误控制程序流程
			if customError == err {
				i = 2
				continue
			}
			t.Error(err)
		}
	}
}

func TestPanicVsExit(t *testing.T) {
	defer func() {
		t.Log("finally")
	}()
	t.Log("start")
	// 对于panic执行会输出详细的错误堆栈信息;并会执行defer
	//panic(errors.New("panic"))
	// 对于os.Exit 不会输出任何堆栈信息且不会执行defer
	os.Exit(-1)
}

func TestRecover(t *testing.T) {
	defer func() {
		// recover 一般应用于应用程序的全局异常处理,而非指定函数内的异常处理
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	t.Log("start")
	//panic(errors.New("异常信息"))
	t.Log("end")
}
