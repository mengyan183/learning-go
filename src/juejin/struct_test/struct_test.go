package struct_test

import (
	"encoding/json"
	"testing"
)

type Person struct {
	age  int
	name string
}

func TestPerson(t *testing.T) {
	// 声明的自定义struct实际是一个成员属性存在默认值的实例
	var p Person
	// struct_test.Person(struct_test.Person{age:0, name:""})
	//assert.Equal(t,nil,p)
	t.Log(p)
	p.age = 1
	p.name = "cs"
	t.Log(p)
	// 这里只是创建了一个指针,并未创建一个空的结构体实例
	ptr := new(Person)
	t.Log(*ptr)
	// 这里默认会对Person进行实例化操作
	ptr = &Person{
		age:  0,
		name: "",
	}
}

func TestNoNameStruct(t *testing.T) {
	s := struct {
		name string
	}{"xing"}
	t.Log(s)
}

type JsonParseObj struct {
	Name   string `json:"name,omitempty"` // json后的字段名称为 name,且当数据为空时不会转化
	Age    int    `json:"-"`              // 代表不会转换该字段
	Height int    `json:"height,string"`  //不管字段类型是什么都以字符串类型返回
}

func TestJsonParse(t *testing.T) {
	s, err := json.Marshal(JsonParseObj{})
	t.Log(string(s), err)
	var obj JsonParseObj
	err = json.Unmarshal(s, &obj)
	t.Log(obj, err)
}
