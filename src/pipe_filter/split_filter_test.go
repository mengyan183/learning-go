package pipe_filter

import "testing"

func TestSplitFilter_Process(t *testing.T) {
	ret, e := NewSplitFilter(",").Process("1,2,3")
	//ret, e := NewSplitFilter(",").Process(1)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(ret)
	}
}
