package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	s, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	sArr := strings.Split(s, sf.delimiter)
	return sArr, nil
}
