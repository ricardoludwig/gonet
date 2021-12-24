package main

import (
	"fmt"
	"github.com/ricardoludwig/gonet/net/http"
)

func main() {
	resp, _ := http.Get{
		Url: "http://www.gmail.com",
	}.Get()
	fmt.Println(resp.Body().ToString())
}
