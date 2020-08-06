package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type ForReflect struct {
	Name string
	Age  int
}

func (fr *ForReflect) ToString() string {
	return fmt.Sprint(fr)
}

func TestValueAndPtr(t *testing.T) {
	fr := ForReflect{
		Name: "xingguo",
		Age:  18,
	}
	t.Logf("%p\n",&fr)
	t.Logf("%p\n",&(fr.Name))
	t.Logf("%p\n",&(fr.Age))
	t.Log(unsafe.Pointer(&(fr.Name)))
	frc := fr
	t.Logf("%p\n",&frc)
	t.Logf("%p\n",&(frc.Name))
	t.Log(unsafe.Pointer(&(frc.Name)))
	t.Log(unsafe.Pointer(&fr))
}

func TestReflectMethod(t *testing.T) {
	fr := ForReflect{
		Name: "xingguo",
		Age:  18,
	}
	// 下面两个语句有什么区别?
	reflect.ValueOf(fr)
	reflect.ValueOf(&fr)
	// reflect.ValueOf(指针/实例)
	s := reflect.ValueOf(&fr).MethodByName("ToString").Call([]reflect.Value{})
	t.Log(s)
}


func TestReflect(t *testing.T) {
	frPtr := new(ForReflect)
	// 获取当前变量的类型
	frPtrT := reflect.TypeOf(frPtr)
	// 如果当前变量的类型为指针
	if frPtrT.Kind() == reflect.Ptr {
		t.Log(frPtrT.Elem().Name())
		// 获取当前指针实际执行元素的类型
		k := frPtrT.Elem().Kind()
		t.Log(k)
		// 如果当前类型为结构体
		if k == reflect.Struct {
			// 获取当前指针执行元素类型的字段数
			nf := frPtrT.Elem().NumField()
			for i := 0; i < nf; i++ {
				// 按照顺序获取具体的结构体字段
				sf := frPtrT.Elem().Field(i)
				t.Log(sf)
			}
		}
	}

	fr := ForReflect{}
	frT := reflect.TypeOf(fr)
	t.Log(frT)
	nf := frT.NumField()
	for i := 0; i < nf; i++ {
		t.Log(frT.Field(i))
	}
}

func TestReflectValue(t *testing.T) {
	// 基本数据类型
	i := 10
	iv := reflect.ValueOf(i)
	// 通过断言转换为指定类型
	t.Log(iv.Int())

	fr := ForReflect{
		Name: "xingguo",
		Age:  18,
	}
	frv := reflect.ValueOf(fr)
	t.Log(frv)
	if frv.Kind() == reflect.Struct {
		t.Log(frv.IsZero(), frv.IsValid())
		if !frv.IsZero() {
			// 获取成员字段的数量
			nf := frv.NumField()
			// 反射获取当前结构体的指针指向的全部元素, 对于结构体实例中存储的元素实际是通过指针
			frvr := reflect.ValueOf(&fr).Elem()
			for i = 0; i < nf; i++ {
				frvrf := frvr.Field(i)
				t.Log(frvrf.Interface())
				// 修改值
				switch frvrf.Kind() {
				case reflect.Int:
					frvrf.Set(reflect.ValueOf(0))
				case reflect.String:
					frvrf.Set(reflect.ValueOf("反射"))
				}
			}
		}
	}
	t.Log(fr)
}

func TestReflectSetValue(t *testing.T) {
	x := 3
	reflect.ValueOf(&x).Elem().SetInt(4)
	t.Log(x)
}
