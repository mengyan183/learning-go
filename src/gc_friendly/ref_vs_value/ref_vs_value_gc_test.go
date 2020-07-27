package ref_vs_value

import "testing"

// 测试 引用传递和值传递 gc效果

const NumsOfElements = 10000

type Content struct {
	Values [NumsOfElements]int
}

// 值传递
func WithValue(contents [NumsOfElements]Content) {

}

// 引用传递
func WithRef(contents *[NumsOfElements]Content) {

}
// 切片内数据也是引用传递
func WithSliceContentRef(contents *[NumsOfElements]*Content) {

}

func TestFunc(t *testing.T) {
	var contents [NumsOfElements]Content
	WithValue(contents)
	WithRef(&contents)
}

func BenchmarkWithValue(b *testing.B) {
	var contents [NumsOfElements]Content

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		WithValue(contents)
	}
	b.StopTimer()
}

func BenchmarkWithRef(b *testing.B) {
	var contents [NumsOfElements]Content

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		WithRef(&contents)
	}
	b.StopTimer()
}

func BenchmarkWithSliceContentRef(b *testing.B) {
	var contents [NumsOfElements]*Content

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		WithSliceContentRef(&contents)
	}
	b.StopTimer()
}
