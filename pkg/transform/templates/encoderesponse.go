package foo

import (
	"context"
	"net/http"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
)

type ResStatus string

var ResponseMessage = map[ResStatus]int{
	Invalid: 400,
}

const (
	Invalid ResStatus = "invalid"
)

func (c ResStatus) String() string {
	return string(c)
}

func (c ResStatus) Error() error {
	return errors.New(string(c))
}

func (c ResStatus) Wrap(err error) error {
	return errors.Wrap(err, string(c))
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"message,omitempty"`
}

type Failure interface {
	Failed() error
}

type Errorer interface {
	Error() error
}

func Error(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}

func JsonError(ctx context.Context, err error, w http.ResponseWriter) {
	headers, ok := ctx.Value("response-headers").(map[string]string)
	if ok {
		for k, v := range headers {
			w.Header().Set(k, v)
		}
	}

	_ = kithttp.EncodeJSONResponse(ctx, w, map[string]interface{}{
		"message": err.Error(),
		"code":    ResponseMessage[ResStatus(strings.Split(err.Error(), ":")[0])],
		"success": false,
	})
}

func JsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(Failure); ok && f.Failed() != nil {
		JsonError(ctx, f.Failed(), w)
		return nil
	}
	resp := response.(Response)
	if resp.Error == nil {
		resp.Code = 200
		resp.Success = true
	}

	headers, ok := ctx.Value("response-headers").(map[string]string)
	if ok {
		for k, v := range headers {
			w.Header().Set(k, v)
		}
	}

	return kithttp.EncodeJSONResponse(ctx, w, resp)
}
