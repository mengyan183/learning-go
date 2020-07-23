package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// 自定义收集器实现类
type CustomCollector struct {
	evtReceiver EventReceiver
	stopChan    chan struct{}
	name        string
	content     string
}

// 自定义实现Collector接口
func (customCollector *CustomCollector) Init(receiver EventReceiver) error {
	customCollector.evtReceiver = receiver
	return nil
}

// 启动start
func (customCollector *CustomCollector) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			customCollector.stopChan <- struct{}{}
			fmt.Print("收到context的关闭消息")
			return nil
		default:
			customCollector.evtReceiver.OnEvent(Event{
				Source:  customCollector.name,
				Content: customCollector.content,
			})
		}
	}
}

func (customCollector *CustomCollector) Stop() error {
	select {
	case <-customCollector.stopChan:
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("current collector not stop")
	}
}
func (customCollector *CustomCollector) Destroy() error {
	return nil
}
func NewCollector(name string, content string) *CustomCollector {
	return &CustomCollector{
		name:    name,
		content: content,
	}
}

func TestAgent_Start(t *testing.T) {
	agent := NewAgent(10)
	for i := 0; i < 2; i++ {
		var build strings.Builder
		build.WriteString("n")
		build.WriteString(strconv.Itoa(i))
		var cbuild strings.Builder
		cbuild.WriteString("c")
		cbuild.WriteString(strconv.Itoa(i))
		collector := NewCollector(build.String(), cbuild.String())
		if err := agent.registerAllCollectors(collector.name, collector); err != nil {
			t.Error(err)
		}
	}
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for {
			t.Log("start agent")
			if err := agent.Start(); err != nil {
				t.Error(err)
				break
			}
		}
		wg.Done()
	}()
	time.Sleep(time.Second * 1)
	go func() {
		wg.Add(1)
		for {
			t.Log("stop agent")
			if err := agent.Stop(); err != nil {
				t.Error(err)
				if _, ok := err.(CollectorsError); !ok {
					break
				}
			}
		}
		wg.Done()
	}()
	time.Sleep(time.Second * 1)
	go func() {
		wg.Add(1)
		for {
			t.Log("destroy agent")
			if err := agent.Destroy(); err == nil {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestContextClose(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				t.Log("收到context的关闭消息")
				return
			default:
				t.Log("当前context尚未关闭")
			}
		}
	}()
	time.Sleep(time.Second * 1)
	go func() {
		cancelFunc()
	}()
	wg.Wait()
}
