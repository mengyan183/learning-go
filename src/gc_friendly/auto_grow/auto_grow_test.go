package auto_grow

import "testing"

const numOfElems = 100000
const times = 1000

// 切片自动扩容
func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		var s []int
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}
// 事先定义好容量的切片,避免自动扩容
func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, numOfElems)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 800000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElems)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElems*8)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

