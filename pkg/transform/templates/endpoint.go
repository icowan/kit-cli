/**
 * @Time : 2020/9/16 3:25 PM
 * @Author : solacowa@gmail.com
 * @File : endpoint
 * @Software: GoLand
 */

package foo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type barRequest struct {
}

type barResponse struct {
}

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
}
