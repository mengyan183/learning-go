package concurrency

import "testing"

func TestV(t *testing.T) {
	t.Logf("%v", int(10))
	t.Logf("%v", []int{1, 2, 3})
	var ch chan int
	t.Logf("%v", ch)
	var ar [1]int
	t.Logf("%v", ar)
	var it interface{}
	t.Logf("%v", it)
}

func TestNoBufferChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		// 异步协程往channel中写入数据
		for i := 0; i < 10; i++ {
			t.Log("准备写入", i)
			ch <- i
			t.Log("写入", i)
		}
		// 写入结束后需要关闭channel,否则消费方会出现死锁
		close(ch)
	}()
	// 消费者,当channel没有关闭且没有数据时当前协程会处于阻塞状态
	t.Log("准备接受消息")
	for v := range ch {
		t.Log("消费", v)
	}
}

func TestBufferChannel(t *testing.T) {
	// 构建buffer chan
	bufferChan := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			t.Log("准备写入", i)
			// 如果buffer 已满时,写入协程会阻塞
			bufferChan <- i
			t.Log("写入", i)
		}
		close(bufferChan)
	}()
	t.Log("准备接受消息")
	// 如果队列中没有消息也会阻塞等待接收
	for v := range bufferChan {
		t.Log("消费", v)
	}
}
