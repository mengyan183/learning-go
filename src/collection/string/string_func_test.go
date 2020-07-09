package string

import (
	"strconv"
	"strings"
	"testing"
)

// 切割字符串
func TestStringSplit(t *testing.T) {
	s := "1,2,3"
	p := strings.Split(s, ",")
	for _, c := range p {
		t.Log(c)
	}
}

// 拼接字符串
func TestStringJoin(t *testing.T) {
	s := []string{"hello", "world"}
	js := strings.Join(s, " ")
	t.Log(js)
}

// 字符串和其他类型的转换
func TestStringCov(t *testing.T) {
	s := strconv.Itoa(1)
	// 执行字符串连接,说明通过 strconv.Itoa转换函数,将int类型数据转换为了字符串
	t.Log("字符串:" + s)
	// 将字符串转换为数值
	//s = "测试" : 测试error
	if i, err := strconv.Atoi(s); err == nil {
		t.Log(1 + i)
	} else {
		t.Log(err)
	}

}
