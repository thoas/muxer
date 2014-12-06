package muxer

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Request struct {
	Request     *http.Request
	Params      map[string]string
	QueryString map[string]string
}

// CurrentRoute allows to retrieve the current
// from a request
func (r *Request) CurrentRoute() *mux.Route {
	return mux.CurrentRoute(r.Request)
}

func (r *Request) GetHost() string {
	return getHost(r.Request)
}

func (r *Request) IsAjax() bool {
	return matchMap(map[string]string{
		"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
	}, r.Request.Header, true)
}
