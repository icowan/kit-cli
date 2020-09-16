/**
 * @Time : 2020/9/16 3:17 PM
 * @Author : solacowa@gmail.com
 * @File : logging
 * @Software: GoLand
 */

package foo

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
}
