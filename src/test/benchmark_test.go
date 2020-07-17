package test

import (
	"bytes"
	"testing"
)

// benchmark性能测试
func BenchmarkString(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)

		}
	}
	b.StopTimer()
}
