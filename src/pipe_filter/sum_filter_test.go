package pipe_filter

import "testing"

func TestSumFilter_Process(t *testing.T) {
	ret, err := NewSplitFilter(",").Process("1,2,3")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(ret)
		intArr, err := NewToIntFilter().Process(ret)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(intArr)
			sum, err := NewSumFilter().Process(intArr)
			if err != nil {
				t.Error(err)
			} else {
				t.Log(sum)
			}
		}
	}
}
