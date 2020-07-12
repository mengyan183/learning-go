package share_memory

import (
	"sync"
	"testing"
	"time"
)

func TestThreadNotSafe(t *testing.T) {
	count := 0
	for i := 0; i < 1000; i++ {
		// 由于共享变量在线程不安全情况下进行操作,因此结果会小于1000
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Log(count)
}

func TestThreadSafe(t *testing.T) {
	var mute sync.Mutex
	count := 0
	for i := 0; i < 1000; i++ {
		go func() {
			// 保证一定会释放锁
			defer func() {
				mute.Unlock()
			}()
			// 执行加锁操作,实现共享变量操作线程安全
			mute.Lock()
			count++
		}()
	}
	time.Sleep(time.Second * 1)
	// 1000
	t.Log(count)
}

func TestWaitGroup(t *testing.T) {
	var mute sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			// 保证一定会释放锁
			defer func() {
				// 如果在加锁未成功情况下释放锁会抛出 fatal error: sync: unlock of unlocked mutex
				mute.Unlock()
				wg.Done()
			}()
			//panic(errors.New("人为异常"))
			// 执行加锁操作,实现共享变量操作线程安全
			mute.Lock()
			count++
		}()
	}
	// 替换主线程阻塞
	//time.Sleep(time.Second * 1)
	// 使用WaitGroup避免了通过time.Sleep阻塞不确定的协程执行时间
	wg.Wait()
	// 1000
	t.Log(count)
}
