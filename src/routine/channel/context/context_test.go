package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case ch := <-ctx.Done():
		fmt.Println(ch, ctx.Err())
		return true
	default:
		return false
	}
}

// 多消费者
func mutiReceiverEmptyStruct(ctx context.Context) {
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			// 可以看到协程的执行时任意选择执行的
			fmt.Printf("go routine:%v \n", i)
			for {
				// 在执行close1操作时, 当前协程正好处于工作中,因此当前协程就接收到了通道中发送的消息
				if isCancelled(ctx) {
					fmt.Println("协程", i, "退出取消监听channel")
					break
				}
			}
		}(i, ctx)
	}
}

func TestTaskCancel(t *testing.T) {
	ctx, cancelFun := context.WithCancel(context.Background())
	mutiReceiverEmptyStruct(ctx)
	// 当前context发出取消时,会将当前context中的所有协程都进行取消
	// 当不调用cancel时,实际ctx.Done返回的channel是一直阻塞的,因为当调用Done时,实际是返回一个新创建的无缓冲的channel,且没有任何消息投递,因此consumer也是一直阻塞的
	// 当调用cancel方法时,实际会调用已生成channel的close方法,因此会广播给所有的consumer
	cancelFun()
	time.Sleep(time.Second * 1)
}

// 测试任务超时
func TestContextDeadLine(t *testing.T) {
	ctx, cancelFun := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancelFun()
	ch := make(chan struct{})
	go func(chan struct{}) chan struct{} {
		//time.Sleep(time.Second * 1)
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
		return ch
	}(ch)
	select {
	case <-ch:
		t.Log("任务执行结束")
	case <-ctx.Done(): // 当deadLine context超时后,会自动触发信道创建和关闭操作
		t.Log("任务执行超时")
	}
	time.Sleep(time.Second * 4)
}
