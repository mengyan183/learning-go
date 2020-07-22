package pipe_filter

type PipeFilter struct {
	Name    string
	Pattern string
}

var filters []Filter

func NewPipeFilter(name string, pattern string) *PipeFilter {
	filters = append(filters, NewSplitFilter(pattern))
	filters = append(filters, NewToIntFilter())
	filters = append(filters, NewSumFilter())
	return &PipeFilter{
		Name:    name,
		Pattern: pattern,
	}
}

func (pipeFilter *PipeFilter) Process(data Request) (Response, error) {
	var ret interface{} = data
	for _, filter := range filters {
		d, err := filter.Process(ret)
		if err != nil {
			return nil, err
		}
		ret = d
	}
	return ret, nil
}
