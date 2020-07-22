package pipe_filter

type Request interface {
}

type Response interface {
}

type Filter interface {
	Process(request Request) (Response, error)
}
