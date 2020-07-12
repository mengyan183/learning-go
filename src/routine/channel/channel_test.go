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
	ch := make(chan string,1)
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
