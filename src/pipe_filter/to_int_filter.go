package pipe_filter

import (
	"errors"
	"strconv"
)

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

func (toIntFilter *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	var re []int
	for _, s := range parts {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		re = append(re, i)
	}
	return re, nil
}
