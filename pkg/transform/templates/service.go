package foo

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Service interface {
	Bar(ctx context.Context) (err error)
}

type service struct {
	logger log.Logger
}

func (s *service) Bar(ctx context.Context) (err error) {
	panic("implement me")
}

func New(logger log.Logger) Service {
	return &service{logger: logger}
}
