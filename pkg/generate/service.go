/**
 * @Time : 2020/8/24 12:18 PM
 * @Author : solacowa@gmail.com
 * @File : generate
 * @Software: GoLand
 */

package generate

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/dave/jennifer/jen"

	"github.com/icowan/kit-cli/pkg/fs"
)

type Service interface {
	CreateFolderStructure(ctx context.Context, path string) error
	AppendMultilineComment(ctx context.Context, src []string) Service
	Raw(ctx context.Context) *jen.Statement
	Commentf(ctx context.Context, format string, a ...interface{}) Service
	AppendInterface(name string, methods []jen.Code) Service
}

type service struct {
	fs     fs.Service
	raw    *jen.Statement
	logger log.Logger
}

func (s *service) Commentf(ctx context.Context, format string, a ...interface{}) Service {
	s.raw.Commentf(format, a...).Line()
	return s
}

func (s *service) AppendInterface(name string, methods []jen.Code) Service {
	s.raw.Type().Id(name).Interface(methods...).Line()
	return s
}

func (s *service) Raw(ctx context.Context) *jen.Statement {
	return s.raw
}

func (s *service) AppendMultilineComment(ctx context.Context, src []string) Service {
	for i, v := range src {
		if i != len(src)-1 {
			s.raw.Comment(v).Line()
			continue
		}
		s.raw.Comment(v)
	}
	return s
}

func (s *service) CreateFolderStructure(ctx context.Context, path string) error {
	e, err := s.fs.Exists(ctx, path)
	if err != nil {
		return err
	}
	if !e {
		_ = level.Info(s.logger).Log("err", fmt.Sprintf("目录不存在 : %s", path))
		return s.fs.MkdirAll(ctx, path)
	}
	return nil
}

func New(logger log.Logger, fs fs.Service) Service {
	return &service{fs: fs, logger: logger}
}
