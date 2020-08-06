package reflect_test

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type ReflectStructPeople struct {
	name string
	age  int
}

func (rsp ReflectStructPeople) ToString() {
	fmt.Printf("{\"name\":%s,\"age\":%d}\n", rsp.name, rsp.age)
}

func (rsp *ReflectStructPeople) SetName(name string) {
	rsp.name = name
}

func TestReflect(t *testing.T) {
	var i int64 = 100
	rt := reflect.TypeOf(i)
	t.Log(rt)
	rv := reflect.ValueOf(i)
	t.Log(rv.Type())
}

func TestSwitchKind(t *testing.T) {
	switchKind(ReflectStructPeople{})
	switchKind(0)
}

func switchKind(i interface{}) {
	switch t := reflect.TypeOf(i).Kind(); t {
	case reflect.Int, reflect.Int8:
		fmt.Println("int")
	case reflect.Struct:
		fmt.Println("struct")
	default:
		fmt.Println(t)
	}
}

// 访问struct 中的成员
func TestReflectStruct(t *testing.T) {
	s := ReflectStructPeople{
		name: "cs",
		age:  10,
	}
	(&s).ToString()
	s.ToString()
	s.SetName("1")
	t.Log(reflect.ValueOf(s).FieldByName("name"))
	reflect.ValueOf(&s).MethodByName("ToString").Call([]reflect.Value{})
	reflect.ValueOf(s).MethodByName("ToString").Call([]reflect.Value{})
	// 由于set方法是指针修改,未执行值复制,因此修改的是相同内存地址的值
	reflect.ValueOf(&s).MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf("修改名称")})
	s.ToString()
}

func TestCompareMapAndSlice(t *testing.T) {
	m1 := map[string]int{"1": 1, "2": 2, "3": 3}
	m2 := map[string]int{"1": 1, "3": 3, "2": 2}
	m3 := map[string]int{"1": 1, "3": 3, "2": 2, "4": 4}
	// 由于map不能直接使用"=="进行比较,可以采用reflect.DeepEqual进行对比两个map
	t.Log(reflect.DeepEqual(m1, m2))
	t.Log(reflect.DeepEqual(m1, m3))

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{1, 2, 3, 4, 5, 6}
	//t.Log(s1 == s2)
	// 对比slice
	t.Log(reflect.DeepEqual(s2, s1))
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestReflectCommon(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := commonSetting(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	reflect.ValueOf(e)
	reflect.TypeOf(e)
}

// 通用配置设置方法;利用反射实现不同类型数据填充
func commonSetting(i interface{}, s map[string]interface{}) error {
	rt := reflect.TypeOf(i)
	// 判断传入的数据是否为指针类型
	if rt.Kind() != reflect.Ptr {
		return errors.New("kind must be PTR")
	}
	if rt.Elem().Kind() != reflect.Struct {
		return errors.New("PTR must point STRUCT")
	}
	if s == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	rv := reflect.ValueOf(i)
	// 遍历settings,要求key和传入数据的struct 中的属性值相同;且value类型相同
	for k, v := range s {
		// 判断当前传入值的类型是否包含当前map中的k
		// 这里fieldByName获取的是structField
		if field, ok = rv.Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			// 这里是获取值的所有元素
			rv.Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestUnsafeTypeForceConvert(t *testing.T) {
	s := "64"
	i := 64
	f := float64(i)
	t.Logf("%T,%f", f, f)
	f = *(*float64)(unsafe.Pointer(&i))
	t.Logf("%T,%f", f, f)
	i = *(*int)(unsafe.Pointer(&s))
	t.Logf("%T,%d", i, i)

	type myInt int

	var j int = 1
	myI := *(*myInt)(unsafe.Pointer(&j))
	// 两种操作等价
	myI = myInt(j)
	t.Log(myI)
}

// 原子操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		var data []int
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	// 初始化共享指针
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			writeDataFn()
			time.Sleep(time.Microsecond * 100)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			readDataFn()
			time.Sleep(time.Microsecond * 100)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestAtomicCache(t *testing.T) {
	var shareCachePTR unsafe.Pointer
	storeCacheFunc := func(v int) {
		ptr := unsafe.Pointer(&v)
		atomic.StorePointer(&shareCachePTR, ptr)
		t.Log("存储数据", ptr, " : ", v)
	}
	getCacheFunc := func() {
		s := atomic.LoadPointer(&shareCachePTR)
		t.Log("读取数据", s, " : ", *(*int)(s))
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				storeCacheFunc(i)
				wg.Done()
			}(i)
		}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				getCacheFunc()
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
}
