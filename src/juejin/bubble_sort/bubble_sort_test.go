package bubble_sort

import (
	"errors"
	"testing"
)

// 冒泡排序

func TestBubbleSort(t *testing.T) {
	values := []int{4, 3, 14, 85, 34, 27, 91, 95, 26, 12, 32}

	if ascValues, err := bubbleSortAsc(values); err != nil {
		t.Error(err)
	} else {
		t.Log(ascValues)
	}
	if descValues, err := bubbleSortDesc(values); err != nil {
		t.Error(err)
	} else {
		t.Log(descValues)
	}
}

func bubbleSortAsc(values []int) ([]int, error) {
	if values == nil {
		return nil, errors.New("数组不能为空")
	}
	// 利用快慢双指针;每一次循环都能保证 i指针之前的数据都已经按照顺序排序完成
	for i := 0; i < len(values)-1; i++ {
		// 快指针
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[i] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values, nil
}
func bubbleSortDesc(values []int) ([]int, error) {
	if values == nil {
		return nil, errors.New("数组不能为空")
	}
	// 利用快慢双指针;每一次循环都能保证 i指针之前的数据都已经按照顺序排序完成
	for i := 0; i < len(values)-1; i++ {
		// 快指针
		for j := i + 1; j < len(values); j++ {
			if values[j] > values[i] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values, nil
}
