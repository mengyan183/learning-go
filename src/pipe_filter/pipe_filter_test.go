package pipe_filter

import "testing"

func TestPipeFilter_Process(t *testing.T) {
	if re, err := NewPipeFilter("p1", ",").Process("1,2,3"); err != nil {
		t.Error(err)
	} else {
		t.Log(re)
	}
}
