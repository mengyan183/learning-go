package channel

import (
	"fmt"
	"testing"
	"time"
)

func doSomeThing() string {
	fmt.Println("do something")
	return "do something done"
}

func doSomeThingElse() string {
	fmt.Println("do something else")
	return "do something else done"
}

func noBufferChannel() chan string {
	// 创建一个无缓冲区的channel
	ch := make(chan string)
	go func() {
		ch <- doSomeThing()
		fmt.Println("channel do something done")
	}()
	return ch
}

func TestNoBufferChannel(t *testing.T) {
	//consumer := noBufferChannel()
	// 如果没有consumer 接收,producer 实际会一直堵塞等待consumer消费,而不会继续往下进行
	noBufferChannel()
	s := doSomeThingElse()
	fmt.Println(s)
	time.Sleep(time.Second * 1)
	//s1 := <-consumer
	//fmt.Println(s, s1)
}

func bufferChannel() chan string {
	// 创建一个缓冲区大小为1的channel
	ch := make(chan string, 1)
	go func() {
		ch <- doSomeThing()
		fmt.Println("channel do something done")
	}()
	return ch
}

func TestBufferChannel(t *testing.T) {
	//consumer := noBufferChannel()
	// 即使没有consumer 接收,producer 也会继续往下进行
	bufferChannel()
	s := doSomeThingElse()
	fmt.Println(s)
	time.Sleep(time.Second * 1)
	//s1 := <-consumer
	//fmt.Println(s, s1)
}

func selectChannel1() chan string {
	// 当不采用buffer channel 时,如果select等待超时,而没有consumer来消费当前channel,会导致当前channel producer一直阻塞
	ch := make(chan string)
	go func() {
		// 模拟channel响应超时
		time.Sleep(time.Second * 2)
		ch <- "selectChannel1"
		// 由于采用阻塞式channel,当select超时选择导致没有consumer消费消息,从而导致当前producer会一直阻塞
		fmt.Println("selectChannel1执行结束")
	}()
	return ch
}

func selectChannel2() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "selectChannel2"
	}()
	return ch
}

// 测试 使用select 实现channel超时控制
func TestSelectOverTime(t *testing.T) {
	select {
	case ch := <-selectChannel1():
		t.Log(ch)
	case ch := <-selectBufferChannel1():
		// 当采用buffer channel,即使select选择超时控制,但buffer channel 在buffer不满的情况下,不会阻塞producer线程
		t.Log(ch)
	case <-time.After(time.Second * 1):
		t.Error("channel已超时")
	}
	time.Sleep(time.Second * 3)
}

func selectBufferChannel1() chan string {
	// 当不采用buffer channel 时,如果select等待超时,而没有consumer来消费当前channel,会导致当前channel producer一直阻塞
	ch := make(chan string, 1)
	go func() {
		// 模拟channel响应超时
		time.Sleep(time.Second * 2)
		ch <- "selectChannel1"
		// 由于采用阻塞式channel,当select超时选择导致没有consumer消费消息,从而导致当前producer会一直阻塞
		fmt.Println("selectChannel1执行结束")
	}()
	return ch
}

/**
如果select中的多个case都可以执行,则会选择任意一个case进行执行
反之:
 当不存在default时,则select会阻塞,直到其中任意channel响应
当存在default,则会直接运行default
*/
// 测试多路选择
func TestSelectMutiChannel(t *testing.T) {
	select {
	case ch := <-selectChannel1():
		t.Log(ch)
	case ch := <-selectChannel2():
		t.Log(ch)
		//default:
		//	t.Error("no compete channel")
	}
}
