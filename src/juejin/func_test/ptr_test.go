package func_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestPtr(t *testing.T) {
	i := 0
	t.Logf("%p", &i)
	t.Log(unsafe.Pointer(&i))

	iPtr := &i
	t.Log(unsafe.Pointer(iPtr))
	t.Log(unsafe.Pointer(&iPtr))
	*iPtr = 100
	t.Log(i)
	iPtrPtr := &iPtr
	t.Log(unsafe.Pointer(iPtrPtr))
	t.Log(unsafe.Pointer(&iPtrPtr))
}

func ptrParam(i *int) {
	*i = 10
}

func TestValueAsPtrParam(t *testing.T) {
	i := 0
	ptrParam(&i)
	assert.Equal(t, 10, i)
}

func TestPanicAndRecover(t *testing.T) {
	// 当使用了recover 可以保证当前协程不会异常终止
	panicAndRecover()
}

func panicAndRecover() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	panic("人工终止")
}
