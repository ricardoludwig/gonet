package http

type Body struct {
	dataByte []byte
	dataText string
}

func (b Body) ToString() string {
	if len(b.dataText) > 1 {
		return b.dataText
	}
	b.dataText = string(b.dataByte)
	return b.dataText
}
