package singletone

import (
	"fmt"
	"sync"
	"testing"
)

// 实现懒汉式单例模式

type ObjectSingleton struct {
	flag bool
}
type EmptyObjectSingleton1 struct {
}
type EmptyObjectSingleton2 struct {
}

var once sync.Once
var object *ObjectSingleton

// 只会执行一次,且指针执行的内存地址都是相同的
func GetSingleToneObject() *ObjectSingleton {
	once.Do(func() {
		object = new(ObjectSingleton)
		object.flag = true
		fmt.Println("执行实例化")
	})
	return object
}

func TestGetSingleToneObject(t *testing.T) {
	var wg sync.WaitGroup
	// 当ObjectSingleton是一个空的struct,所以每次得到的实例地址都是相同的地址
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingleToneObject()
			//t.Logf("%X", unsafe.Pointer(obj))
			t.Logf("%p", obj)
			wg.Done()
		}()
	}
	wg.Wait()
	// 空的struct对象
	t.Log(struct{}{} == struct{}{})
	t.Logf("%p", &EmptyObjectSingleton1{})
	t.Logf("%p", &EmptyObjectSingleton1{})
	t.Logf("%p", &EmptyObjectSingleton2{})
}
