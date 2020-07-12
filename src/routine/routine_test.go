package routine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestFirstGoRoutine(t *testing.T) {
	// 自定义设置运行时最大处理器个数
	//runtime.GOMAXPROCS(1)
	i := 0
	for ; i < 10; i++ {
		// 启动协程,使用值复制,避免内存共享导致异常
		go func(j int) {
			t.Log(j)
		}(i)
		// 共享内存情况
		go func() {
			//这里输出i全部为10的原因在于, 由于这个没有采用值复制,而是采用共享内存的方式,所有添加到routine队列中的routine中的数据地址都是指向i;
			//因此在for循环结束后,协程还并没有被真正执行调度,当协程真正执行时,i变成了10,所有输出的都是10
			t.Log(i)
		}()
		// 如果每一次循环都阻塞一下当前主线程,给协程调度充足的时间,则输出的结果就会按照i递增结果进行输出
		//time.Sleep(time.Second * 1)
	}
	t.Log("循环执行结束:", i)
	time.Sleep(time.Second * 1)
}
// 测试调度器
func TestSched(t *testing.T)  {
	// 设置运行时的最大processor个数
	runtime.GOMAXPROCS(1)
	go func() {
		for true {
			fmt.Println("1")
		}
	}()
	for true {
		fmt.Println("0")
	}
	// 根据输出的结果可以看到存在协程切换
}
