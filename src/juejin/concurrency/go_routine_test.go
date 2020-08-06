package concurrency

import (
	"runtime"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			t.Log("异步协程", i)
		}
	}() // 对于匿名函数数据传递
	for i := 0; i < 10; i++ {
		t.Log("主协程")
	}
	t.Log("主协程休眠")
	time.Sleep(time.Microsecond)
	t.Log("主协程休眠结束")
}
func TestGoRoutineWithCustom(t *testing.T) {
	// 通过获取当前逻辑cpu数量设置最大可运行processor数量 ;  一般更推荐使用在init方法中设置
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func() {
		defer func() {
			t.Log("defer")
		}()
		for i := 0; i < 10; i++ {
			t.Log("异步协程", i)
			if i == 5 {
				// 中止并销毁当前协程,其他协程不会受到影响,且如果存在defer会将defer执行结束后再完全退出
				runtime.Goexit() // 不可使用在主协程(main函数)中会导致panic
			}
		}
	}() // 对于匿名函数数据传递
	for i := 0; i < 10; i++ {
		t.Log("主协程")
	}
	t.Log("主协程休眠")
	time.Sleep(time.Microsecond)
	t.Log("主协程休眠结束")
}
