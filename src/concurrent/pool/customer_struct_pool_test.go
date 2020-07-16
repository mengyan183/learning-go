package pool

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"unsafe"
)

func TestEmptyCustomerStructPool(t *testing.T) {
	runtime.GOMAXPROCS(10)
	pool := new(EmptyCustomerStructPool)
	noEmptypool := new(CustomerStructPool)
	for i := 0; i < 10; i++ {
		go func() {
			if s, e := pool.Get(time.Second * 1); e != nil {
				taskEmpty(s)
				if e := pool.Release(s); e != nil {
					t.Error(e)
				}
			} else {
				t.Error(e)
			}
		}()
		go func() {
			if s, e := noEmptypool.Get(time.Second * 1); e != nil {
				task(s)
				if e := noEmptypool.Release(s); e != nil {
					t.Error(e)
				}
			} else {
				t.Error(e)
			}
		}()
	}

	time.Sleep(time.Second * 10)
}

func taskEmpty(e *EmptyCustomerStruct) {
	fmt.Printf("%X,空对象执行任务\n", unsafe.Pointer(&e))
}
func task(e *CustomerStruct) {
	fmt.Printf("%X,执行任务\n", unsafe.Pointer(&e))
}
