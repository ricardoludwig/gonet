package http

type Status struct {
	code int
	msg  string
}

func (s Status) Code() int {
	return s.code
}

func (s Status) Message() string {
	return s.msg
}
