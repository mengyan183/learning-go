package close

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 发布者
func publisher(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
}

// 发布者自定义关闭channel
func publisherWithClose(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭channel
		close(ch)
	}()
}

func receiverWithCloseMessage(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 当采用死循环一直去接受channel时,当channel中没有数据时会抛出 deadLock
		for {
			//for i := 0; i < 10; i++ {
			if v, ok := <-ch; ok {
				fmt.Println(v, ok)
			} else {
				fmt.Println("channel已关闭", v, ok)
				break
			}
		}
		wg.Done()
	}()
}

// 接受者
//func receiver(ch chan int, wg sync.WaitGroup) { 当采用值接收时,
func receiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 当采用死循环一直去接受channel时,当channel中没有数据时会抛出 deadLock
		for {
			//for i := 0; i < 10; i++ {
			if v, ok := <-ch; ok {
				fmt.Println(v, ok)
			} else {
				fmt.Println(v, ok)
			}
		}
		wg.Done()
	}()
}

func TestPubRe(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	publisher(ch)
	wg.Add(1)
	//receiver(ch, wg)
	receiver(ch, &wg)
	wg.Add(1)
	//receiver(ch, wg)
	receiver(ch, &wg)
	wg.Wait()
}

func TestChannelWithBuffer(t *testing.T) {
	var wg sync.WaitGroup
	// 实例化有buffer的channel
	ch := make(chan int, 1)
	publisher(ch)
	wg.Add(1)
	receiver(ch, &wg)
	wg.Wait()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	publisherWithClose(ch)
	wg.Add(1)
	receiverWithCloseMessage(ch, &wg)
	wg.Add(1)
	receiverWithCloseMessage(ch, &wg)
	wg.Wait()
}

// 通过手动发送取消消息实现自定义关闭
func closeChannel1(ch chan struct{}) {
	time.Sleep(time.Microsecond * 1)
	ch <- struct{}{}
}
func closeChannel2(ch chan struct{}) {
	close(ch)
}
func isCancelled(ch chan struct{}) bool {
	select {
	// 这里也是作为receiver来接收channel的消息来判断当前channel是否关闭
	case <-ch:
		return true
	default:
		return false
	}
}

// 多消费者
func mutiReceiverEmptyStruct(ch chan struct{}) {
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			// 可以看到协程的执行时任意选择执行的
			fmt.Printf("go routine:%v \n", i)
			for {
				// 在执行close1操作时, 当前协程正好处于工作中,因此当前协程就接收到了通道中发送的消息
				if isCancelled(ch) {
					fmt.Println("协程", i, "退出取消监听channel")
					break
				}
			}
		}(i, ch)
	}
	// 采用第一种方式取
	closeChannel1(ch)
	// 当采用广播的形式来执行channel关闭时,所有的receiver都会接收到当前消息
	//closeChannel2(ch)
	time.Sleep(time.Second * 1)
}

func TestTaskCancel(t *testing.T) {
	//mutiReceiverEmptyStruct(make(chan struct{}))
	//mutiReceiverEmptyStruct(make(chan struct{},10))
}
