package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler func(Response, *Request)

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	queryString := make(map[string]string)

	if len(req.Form) > 0 {
		for k, v := range req.Form {
			queryString[k] = v[0]
		}
	}

	h(Response{
		w,
	}, &Request{
		Request:     req,
		QueryString: queryString,
		Params:      mux.Vars(req),
	})
}
