/**
 * @Time : 2020/8/24 12:22 PM
 * @Author : solacowa@gmail.com
 * @File : service
 * @Software: GoLand
 */

package service

import (
	"context"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/go-kit/kit/log"
	"github.com/icowan/kit-cli/pkg/fs"
	"github.com/icowan/kit-cli/pkg/generate"
	"github.com/kujtimiihoxha/kit/utils"
	"path"
	"strings"
)

type Service interface {
	Gen(ctx context.Context) error
}

type service struct {
	fs            fs.Service
	gen           generate.Service
	interfaceName string
	name          string
	filePath      string
	srcFile       *jen.File
	logger        log.Logger
}

func (s *service) Gen(ctx context.Context) (err error) {
	if err = s.gen.CreateFolderStructure(ctx, fmt.Sprintf("./pkg/%s", s.name)); err != nil {
		fmt.Println(fmt.Sprintf("gen.CreateFolderStructure err: %v", err))
		return err
	}

	comments := []string{
		"将您的方法添加到这里",
		"e.x: Foo(ctx context.Context,s string)(rs string, err error)",
	}

	s.gen.AppendMultilineComment(ctx, comments).
		Raw(ctx).Commentf("%s describes the service.", s.interfaceName).Line()

	s.gen.AppendInterface(
		s.interfaceName,
		[]jen.Code{s.gen.Raw(ctx)})

	return s.fs.WriteFile(ctx, s.filePath, s.srcFile.GoString(), false)
}

func New(logger log.Logger, name, servicePathFormat string, gen generate.Service) Service {
	destPath := fmt.Sprintf(servicePathFormat, utils.ToLowerSnakeCase(name))
	return &service{
		name:          name,
		interfaceName: "Service",
		filePath:      path.Join(destPath, "service.go"),
		srcFile:       jen.NewFilePath(strings.Replace(destPath, "\\", "/", -1)),
		gen:           gen,
		logger:        logger,
	}
}
