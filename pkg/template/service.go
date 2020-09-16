/**
 * @Time : 2020/9/16 10:50 AM
 * @Author : solacowa@gmail.com
 * @File : service
 * @Software: GoLand
 */

package template

import "golang.org/x/tools/godoc/vfs/mapfs"

var ASTTemplates = mapfs.New(map[string]string{
	`service.go`: "package foo\n\nimport (\n	\"context\"\n	\"github.com/go-kit/kit/log\"\n)\n\ntype Service interface {\n	Bar(ctx context.Context) (err error)\n}\n\ntype service struct {\n	logger log.Logger\n}\n\nfunc (s *service) Bar(ctx context.Context) (err error) {\n	panic(\"implement me\")\n}\n\nfunc New(logger log.Logger) Service {\n	return &service{logger: logger}\n}",
	`endpoint.go`: `package foo
import (
	"context"
	"github.com/go-kit/kit/endpoint"
)
type barRequest struct {}
type barResponse struct {}
type Endpoints struct {
	BarEndpoint endpoint.Endpoint
}
func NewEndpoint(s Service, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		BarEndpoint: makeBarEndpoint(s),
	}
	for _, m := range mdw["Bar"] {
		eps.BarEndpoint = m(eps.BarEndpoint)
	}
	return eps
}
func makeBarEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		err = s.Bar(ctx)
		return barResponse{}, err
	}
}`,
	`http.go`: `package foo
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
}`,
	`grpc.go`: "package foo\n",
	`logging.go`: `package foo
import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)
type loggingServer struct {
	logger log.Logger
	Service
}
func NewLoggingServer(logger log.Logger, s Service) Service {
	return &loggingServer{
		logger:  level.Info(logger),
		Service: s,
	}
}
func (s *loggingServer) Bar(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"trace-id", ctx.Value("trace-id"),
			"method", "Bar",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.Bar(ctx)
}`,
	`middleware.go`: "package foo",
})
