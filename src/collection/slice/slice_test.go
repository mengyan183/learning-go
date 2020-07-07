package slice

import "testing"

func TestSlice(t *testing.T) {
	// 声明一个slice,和数组声明的区别,不需要指定长度
	var slice []int
	t.Log(len(slice), cap(slice)) // 0 0
	// 往slice中添加值
	slice = append(slice, 1)
	t.Log(len(slice), cap(slice)) // 1 1
	// 声明并赋值操作,则当前slice的容量和长度都为值的个数
	slice1 := []int{1, 2, 3, 4}
	t.Log(len(slice1), cap(slice1))
	// 声明指定容量和长度的slice; make(结构类型,长度,容量)
	slice2 := make([]int, 3, 5)
	t.Log(len(slice2), cap(slice2))
	t.Log(slice2) //[0 0 0] 如果指定了长度,则按照指定的长度会自动默认初始化数据,且不可访问超过指定长度的索引位置,未初始化元素不可访问
	//t.Log(slice2[3]) // 会抛出越界异常  runtime error: index out of range [3] with length 3
}

func TestSliceAtoGrowing(t *testing.T) {
	// 定义切片
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		// cap 是按照 2的阶乘进行扩容
		t.Log(len(s), cap(s))
	}
}

// 切片共享内存空间
func TestSliceShareMemory(t *testing.T) {
	s := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"}
	// 截取,对于q2来说共享了s已存在的内存空间,并没有生成新的内存空间; 对于q2的内存空间而言实际是从 索引3开始到s的结束位置,从索引3开始的连续内存空间;
	// 所以对于 q2而言,cap容量实际等于 len(s) - 起始截取索引(3);
	q2 := s[3:6]
	t.Log(q2, len(q2), cap(q2))
	summer := s[5:9]
	// len(summer) == 4; cap(summer) == 7
	t.Log(summer, len(summer), cap(summer))
	// 由于切片截取后的切片仍然共享的是同一块内存空间,因此对相同空间的数据进行修改,也会影响到其他的切片
	summer[0] = "Unknown"
	t.Log(q2)
	t.Log(s)
	// 测试截取后的切片,如果发生自动扩容后会对源切片造成什么影响
	end := s[11:]
	t.Log(end, len(end), cap(end))
	end = append(end, "append")
	t.Log(end, len(end), cap(end))
}

// 切片数据对比
func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	//t.Log(a == b)//编译错误:slice can only be compared to nil
	t.Log(a, b)
}

// 对比几种切片初始化方式不同
func TestSliceInitDiff(t *testing.T) {
	var a []int
	b := []int{}
	t.Log(a, a == nil, b, b == nil)
	a = append(a,1)
	t.Log(a)
}
