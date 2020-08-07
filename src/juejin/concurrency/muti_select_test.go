package concurrency

import (
	"testing"
	"time"
)

// 使用select实现多路复用

func TestMutiChanSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for true {
			select {
			case v := <-ch1:
				t.Log(v)
			case ch2 <- 1:
				t.Log("写入数据")
			case <-time.After(time.Second):
				t.Log("已超时")
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			t.Log("ch1写入数据：", i)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			str := <-ch2
			t.Log("获取到ch2的数据：", str)
		}
	}()
	time.Sleep(time.Minute)
}
