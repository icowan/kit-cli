/**
 * @Time : 2020/8/24 12:33 PM
 * @Author : solacowa@gmail.com
 * @File : service
 * @Software: GoLand
 */

package fs

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"

	"github.com/Songmu/prompter"
	"github.com/spf13/afero"
)

type Service interface {
	Exists(ctx context.Context, path string) (bool, error)
	ReadFile(ctx context.Context, path string) (string, error)
	WriteFile(ctx context.Context, path, data string, force bool) error
	Mkdir(ctx context.Context, dir string) error
	MkdirAll(ctx context.Context, path string) error
}

type service struct {
	fs            afero.Fs
	forceOverride bool
	logger        log.Logger
}

func (s *service) Mkdir(ctx context.Context, dir string) error {
	return s.fs.Mkdir(dir, os.ModePerm)
}

func (s *service) MkdirAll(ctx context.Context, path string) error {
	return s.fs.MkdirAll(path, os.ModePerm)
}

func (s *service) WriteFile(ctx context.Context, path, data string, force bool) error {
	if b, _ := s.Exists(ctx, path); b && !(s.forceOverride || force) {
		ss, _ := s.ReadFile(ctx, path)
		if ss == data {
			_ = level.Warn(s.logger).Log("msg", fmt.Sprintf("`%s` 存在且相同，将被忽略", path))
			return nil
		}
		b := prompter.YN(fmt.Sprintf("`%s` 已经存在，是否要覆盖它?", path), false)
		if !b {
			return nil
		}
	}
	return afero.WriteFile(s.fs, path, []byte(data), os.ModePerm)
}

func (s *service) ReadFile(ctx context.Context, path string) (string, error) {
	d, err := afero.ReadFile(s.fs, path)
	return string(d), err
}

func (s *service) Exists(ctx context.Context, path string) (bool, error) {
	return afero.Exists(s.fs, path)
}

func New(logger log.Logger, folder, dir string, forceOverride bool) Service {
	var inFs afero.Fs

	if folder != "" {
		inFs = afero.NewBasePathFs(afero.NewOsFs(), folder)
	} else {
		inFs = afero.NewOsFs()
	}

	var fs afero.Fs

	if dir != "" {
		fs = afero.NewBasePathFs(inFs, dir)
	} else {
		fs = inFs
	}
	return &service{
		fs:            fs,
		forceOverride: forceOverride,
		logger:        logger,
	}
}
