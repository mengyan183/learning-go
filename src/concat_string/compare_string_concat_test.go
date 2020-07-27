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

func BenchmarkStringsBuilder(b *testing.B)  {

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
