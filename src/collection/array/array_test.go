package array

import "testing"

func TestArrayInit(t *testing.T) {
	// 声明数组,默认值都为 元素类型的默认值
	var arr [5]int
	// 赋值
	arr[0] = 1
	//声明且指定数组长度以及初始化赋值
	arr1 := [2]int{1, 2}
	// 声明不指定数组长度以及初始化赋值, 数组的长度为赋值的元素个数
	arr2 := [...]int{1, 2, 3, 4}
	t.Log(len(arr), arr, len(arr1), arr1, len(arr2), arr2)
}

func TestArrayIterator(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	// 迭代器
	for idx, data := range arr {
		t.Log(idx, data)
	}
	// 由于使用range返回两个字段,可以使用占位符的形式来接收不需要的字段
	for _, data := range arr {
		t.Log(data)
	}
	t.Log(arr[:])
}

//数组截取
func TestArraySelector(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	// 截取整个数组
	arrSelector := arr[:]
	t.Log(arrSelector)
	// 从指定索引位置开始截取到结尾
	arrSelector = arr[1:]
	t.Log(arrSelector)
	// 从0开始截取到指定位置的数组
	arrSelector = arr[:1]
	t.Log(arrSelector)
	// 截取指定开始索引位置以及指定结束索引位置
	arrSelector = arr[1:2]
	t.Log(arrSelector)
	// 数组越界
	//arrSelector = arr[:len(arr)+1] //invalid slice index len(arr) + 1 (out of bounds for 5-element array) 编译错误,直接提示数组越界
	//t.Log(arrSelector)
	// 如果结束索引等于开始索引,截取整个数组
	arrSelector = arr[2:2]
	t.Log(arrSelector)
	// 如果结束索引小于开始索引,编译错误
	//arrSelector = arr[2:1]
	//t.Log(arrSelector)
}
