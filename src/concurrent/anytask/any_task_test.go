package anytask

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func AnyTask(numGoRoutine int, ch chan string) string {
	for i := 0; i < numGoRoutine; i++ {
		go func(i int) {
			time.Sleep(10 * time.Millisecond)
			ch <- fmt.Sprintf("The result is from %d", i)
		}(i)
	}
	return <-ch
}

func TestAnyTask(t *testing.T) {
	// 获取当前的协程数量
	t.Log("Before", runtime.NumGoroutine())
	numGoRoutine := 10
	// 当使用阻塞channel时,由于只有一个consumer去消费,会导致其他协程生产者会阻塞到当前channel
	//ch := make(chan string)
	// 当使用buffer channel时,只要channel中存在剩余空间,producer就不会阻塞,协程就会释放当前资源,避免协程泄露
	ch := make(chan string, numGoRoutine)
	t.Log(AnyTask(numGoRoutine, ch))
	// TODO 如果优雅的关闭不需要的channel
	//close(ch)
	time.Sleep(time.Second * 1)
	t.Log("After", runtime.NumGoroutine())

}
