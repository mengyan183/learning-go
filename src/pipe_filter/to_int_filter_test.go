package pipe_filter

import "testing"

func TestToIntFilter_Process(t *testing.T) {
	ret, e := NewSplitFilter(",").Process("1,2,3")
	//ret, e := NewSplitFilter(",").Process("1,a,3")
	if e != nil {
		t.Error(e)
	} else {
		t.Log(ret)
		intArr, err := NewToIntFilter().Process(ret)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(intArr)
		}
	}
}
