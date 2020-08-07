package concurrency

import (
	"sync"
	"testing"
)

func TestDeadLock(t *testing.T) {
	c := make(chan int)
	c <- 88
	t.Log("不可达代码")
	// 在协程启动之前往channel中写入数据导致主协程阻塞
	go func() {
		<-c
	}()
}

func TestDeadLockWithWG(t *testing.T) {
	var wg sync.WaitGroup
	// 由于等待同步数量大于实际完成(Done)数量,导致协程一直Wait 因此死锁
	wg.Add(1)
	wg.Wait()
}

// 在同一个协程内对同一个channel同时执行读写操作
func TestDeadLockWithChannel(t *testing.T) {
	ch := make(chan int)
	ch <- 8
	// 这里是不可达代码
	<-ch
}

func TestMutiChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			c := <-ch1
			ch2 <- c
		}
	}()

	for {
		c := <-ch2
		ch1 <- c
	}
}
