package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	Waiting = iota
	Running
)

type CollectorsError struct {
	CollectorErrors []error
}

func (ce CollectorsError) Error() string {
	var strs []string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

// 事件
type Event struct {
	// 来源
	Source string
	// 内容
	Content string
}

// 事件接收器
type EventReceiver interface {
	// 处理事件
	OnEvent(event Event)
}

// 收集器接口
type Collector interface {
	// 初始化
	Init(receiver EventReceiver) error
	// 启动, 为了保证收集到的事件都在不同的协程中进行处理,且对协程进行统一管理
	Start(ctx context.Context) error
	Stop() error
	Destroy() error
}

// 插件
type Agent struct {
	// 所有的收集器
	collectors map[string]Collector
	// 事件channel
	evtBuf chan Event
	// 域取消方法
	cancel context.CancelFunc
	// 域
	ctx context.Context
	// 状态; 状态机的效果,为了避免当前agent重复启动
	state int
}

// 重写OnEvent方法
func (agent *Agent) OnEvent(event Event) {
	// 将event写入到channel中
	agent.evtBuf <- event
}

// 注册register
func (agent *Agent) registerAllCollectors(name string, collector Collector) error {
	if agent.state != Waiting {
		return errors.New("agent has started")
	}
	if name == "" || collector == nil {
		return errors.New("argu cannot be nil")
	}
	agent.collectors[name] = collector
	// 初始化
	return collector.Init(agent)
}

// 启动所有的收集器
func (agent *Agent) startAllCollectors() error {
	var err error
	var errs CollectorsError
	var mutx sync.Mutex
	// 保证所有的collector都已启动
	var wg sync.WaitGroup
	if len(agent.collectors) == 0 {
		return nil
	}
	for name, collector := range agent.collectors {
		wg.Add(1)
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutx.Unlock()
				wg.Done()
			}()
			// 使用context来统一管理当前agent中所有collector的协程
			err = collector.Start(ctx)
			mutx.Lock()
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agent.ctx)
	}
	wg.Wait()
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

// 关闭所有的collector
func (agent *Agent) stopAllCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agent.collectors {
		err = collector.Stop()
		if err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

// 销毁所有的collector
func (agent *Agent) destroyAllCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agent.collectors {
		err = collector.Destroy()
		if err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

// 接收当前agent中的event
func (agent *Agent) eventChannelGoRoutine() {
	for true {
		select {
		case event := <-agent.evtBuf:
			fmt.Printf("source:%s,content:%s\n", event.Source, event.Content)
		case <-agent.ctx.Done():
			return
		}
	}
}

// 启动当前agent
func (agent *Agent) Start() error {
	if agent.state != Waiting {
		return errors.New("agent has started")
	}
	// 设置当前agent的状态为运行中
	agent.state = Running
	// 创建当前agent的context以及cancel方法
	agent.ctx, agent.cancel = context.WithCancel(context.Background())
	// 启动当前接收事件的异步线程
	go agent.eventChannelGoRoutine()
	return agent.startAllCollectors()
}

// 关闭当前agent
func (agent *Agent) Stop() error {
	if agent.state != Running {
		return errors.New("agent has stoped")
	}
	// 设置当前agent 为 wait
	agent.state = Waiting
	// 关闭当前agent中的context
	agent.cancel()
	return agent.stopAllCollectors()
}

// 回收资源
func (agent *Agent) Destroy() error {
	if agent.state != Waiting {
		return errors.New("agent has started")
	}
	return agent.destroyAllCollectors()
}

func NewAgent(bufLength int) *Agent {
	return &Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, bufLength),
		state:      Waiting,
	}
}
