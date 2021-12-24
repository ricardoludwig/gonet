package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Get struct {
	Url    string
	resp   Response
	client http.Client
}

func (g Get) Get() (resp Response, err error) {

	rsp, err := g.client.Get(g.Url)
	if rsp != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("There is something wrong, failed to close the connection. ", err)
				return
			}
		}(rsp.Body)
	}

	if err != nil {
		log.Println("There is something wrong. Connection fail, check service availability or URL. ", err)
		return
	}

	bd, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println("Fail, response body unavailable.", err)
		return
	}

	return Response{
		bd: Body{
			dataByte: bd,
		},
		st: Status{
			code: rsp.StatusCode,
			msg:  rsp.Status,
		},
	}, nil

}
