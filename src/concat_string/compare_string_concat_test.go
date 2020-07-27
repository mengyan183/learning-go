package concat_string

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100

func BenchmarkFmtSprintf(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStrConv(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}

func BenchmarkBytesBuffer(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for i := 0; i < numbers; i++ {
			buffer.WriteString(strconv.Itoa(i))
		}
		_ = buffer.String()
	}
	b.StopTimer()
}

func BenchmarkStringsBuilder(b *testing.B) {

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func TestMapBasicDataType(t *testing.T) {
	m := map[string]string{"1": "1"}
	editMap(m)
	// change
	t.Log(m["1"])
	im := map[string]int{"1": 1}
	editIntMap(im)
	// 2
	t.Log(im["1"])
}
func editIntMap(m map[string]int) {
	if m == nil {
		return
	}
	m["1"] = 2
}
func editMap(m map[string]string) {
	if m == nil {
		return
	}
	m["1"] = "change"
}

type T struct {
	i int
}

// 对比值和引用传递
func TestMapStructType(t *testing.T) {
	m := map[string]T{"1": {1}}
	editStructMap(m)
	// 1
	t.Log(m["1"].i)
	// 从这里可以看出来,当map存储的是struct实例时,实际值复制
	v, _ := m["1"]
	v.i = 3
	t.Log(m["1"].i)
	t.Log(v)
	refm := map[string]*T{"1": {1}}
	editStructRefMap(refm)
	// 2
	t.Log(refm["1"].i)
}

func editStructRefMap(m map[string]*T) {
	if m == nil {
		return
	}
	v, _ := m["1"]
	v.i = 2
}

func editStructMap(m map[string]T) {
	if m == nil {
		return
	}
	v, _ := m["1"]
	v.i = 2
}
