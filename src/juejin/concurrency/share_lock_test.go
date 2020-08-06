package concurrency

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 对于有限共享资源的情况下,一般推荐使用WaitGroup避免过量访问

var wg sync.WaitGroup
var total = 2

func TestShareAndWg(t *testing.T) {
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			t.Log("执行")
			total--
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestWithLock(t *testing.T) {
	var m sync.Mutex
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			defer func() {
				m.Unlock()
			}()
			m.Lock()
			t.Log("执行")
			total--
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestWithRwLock(t *testing.T) {
	var rw sync.RWMutex
	for i := 0; i < 100; i++ {
		go func() {
			readWithLock(&rw)
		}()
		go func() {
			writeWithLock(&rw)
		}()
	}
	time.Sleep(time.Minute * 5)
}

func readWithLock(rw *sync.RWMutex) {
	defer func() {
		rw.RUnlock()
		fmt.Println("结束读")
	}()
	rw.RLock()
	fmt.Println(goID(),"开始读")
	time.Sleep(time.Second)
}

func writeWithLock(rw *sync.RWMutex) {
	defer func() {
		rw.Unlock()
		fmt.Println("结束写")
	}()
	rw.Lock()
	fmt.Println(goID(),"开始写")
	time.Sleep(time.Second)
}

func goID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
