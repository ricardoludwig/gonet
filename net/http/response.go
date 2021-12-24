package http

type Response struct {
	bd Body
	st Status
}

func (r Response) Body() Body {
	return r.bd
}

func (r Response) Status() Status {
	return r.st
}

func (r Response) StatusCode() int {
	return r.st.Code()
}

func (r Response) StatusMessageCode() string {
	return r.st.Message()
}
