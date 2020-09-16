/**
 * @Time : 2020/9/16 3:25 PM
 * @Author : solacowa@gmail.com
 * @File : http
 * @Software: GoLand
 */

package foo

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	rateBucketNum = 100
)

func MakeHTTPHandler(logger kitlog.Logger, s Service, opts []kithttp.ServerOption) http.Handler {
	var ems []endpoint.Middleware

	s = NewLoggingServer(logger, s)

	eps := NewEndpoint(s, map[string][]endpoint.Middleware{
		"Bar": ems,
	})

	r := mux.NewRouter()
	r.Handle("/foo/bar", kithttp.NewServer(
		eps.BarEndpoint,
		decodeBarRequest,
		encode.ResponseJson,
		opts...,
	)).Methods(http.MethodGet)
	return r

}

func decodeBarRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return barRequest{}, nil
}
