package extension

import (
	"fmt"
	"testing"
)

// 继承测试

type Human struct {
	Age  int
	Name string
}

func (human *Human) Speak(word string) {
	fmt.Printf("human %s 说 %s\n", human.Name, word)
}

func (human *Human) SpeakTo(word string) {
	human.Speak(word)
}

type Man struct {
	// 匿名内置类
	*Human
}

// 重写Man中的方法
func (man *Man) Speak(word string) {
	fmt.Printf("man %s 说 %s\n", man.Name, word)
}

func TestAnonymousInnerStruct(t *testing.T) {
	// 由于go不存在继承,因此实例化的Man不能转换为Human类型
	// 编译报错:cannot use Man literal (type Man) as type Human in assignment
	//var human Human = Man{Human{18,"man"}}

	var man Man = Man{&Human{18, "man"}}
	man.Speak("HelloWorld")
	// 这里调用的是human中的Speak方法,并没有调用man中的speak方法;并没有达到多态的效果
	man.SpeakTo("HelloWorld")
}
