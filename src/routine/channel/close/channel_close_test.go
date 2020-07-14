package close

import (
	"fmt"
	"sync"
	"testing"
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
				fmt.Println(v,ok)
			} else {
				fmt.Println("channel已关闭",v,ok)
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
				fmt.Println(v,ok)
			} else {
				fmt.Println(v,ok)
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
