package pipe_filter

import "errors"

type SumFilter struct {
}

var SumFilterWrongFormatError = errors.New("input data should be []int")

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sumFilter *SumFilter) Process(request Request) (Response, error) {
	intArr, ok := request.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	var re int
	for _, v := range intArr {
		re += v
	}
	return re, nil
}
