package cache

import (
	"sync"
	"testing"
)

// 对象缓存测试

func TestSyncPool(t *testing.T) {
	// 创建 缓存
	cache := &sync.Pool{
		// 如果缓存中都不存在,则调用当前方法返回新的数据
		New: func() interface{} {
			t.Log("create new Object")
			return 100
		},
	}
	for i := 0; i < 2; i++ {
		v := cache.Get().(int)
		// 可以看出如果缓存中没有数据,创建新的对象不会自动写入到缓存中
		t.Log("缓存中没有数据,第一次获取", v)
	}
	// 往缓存中写入数据,当私有对象中不存在数据时,会将数据写入到私有对象区;
	//而对于缓存区的数据结构为栈,规则为后进先出,因此可以看到从缓冲区获取数据的顺序和写入的顺序相反
	for i := 0; i < 4; i++ {
		cache.Put(i)
	}
	// gc,
	//runtime.GC()
	//time.Sleep(time.Second * 1)
	for i := 0; i < 5; i++ {
		t.Log(cache.Get().(int))
	}
}
