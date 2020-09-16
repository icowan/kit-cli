/**
 * @Time : 2020/9/16 10:46 AM
 * @Author : solacowa@gmail.com
 * @File : service
 * @Software: GoLand
 */

package transform

import (
	"bytes"
	"context"
	"fmt"
	"github.com/icowan/kit-cli/pkg/template"
	"github.com/pkg/errors"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/tools/imports"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func importPath(targetDir, gopath string) (string, error) {
	if !filepath.IsAbs(targetDir) {
		return "", fmt.Errorf("%q is not an absolute path", targetDir)
	}

	for _, dir := range filepath.SplitList(gopath) {
		abspath, err := filepath.Abs(dir)
		if err != nil {
			continue
		}
		srcPath := filepath.Join(abspath, "src")

		res, err := filepath.Rel(srcPath, targetDir)
		if err != nil {
			continue
		}
		if strings.Index(res, "..") == -1 {
			return res, nil
		}
	}
	return "", fmt.Errorf("%q is not in GOPATH (%s)", targetDir, gopath)

}

type Service interface {
	TransformAST(ctx context.Context, pkgName string, transport ...string)
	Init(ctx context.Context, name string) (err error)
}

type service struct {
}

var genFiles = []string{
	"service.go",
	"endpoint.go",
	"logging.go",
	"http.go",
	"grpc.go",
	"middleware.go",
}

func getGopath() string {
	gopath, set := os.LookupEnv("GOPATH")
	if !set {
		return filepath.Join(os.Getenv("HOME"), "go")
	}
	return gopath
}

func (s *service) Init(ctx context.Context, name string) error {
	gopath := getGopath()
	dir, _ := os.Getwd()
	importBase, err := importPath(fmt.Sprintf(dir+"/world"), gopath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(importBase)

	return nil
}

func (s *service) TransformAST(ctx context.Context, pkgName string, transport ...string) {
	for _, v := range genFiles {
		fileName := fmt.Sprintf("src/pkg/%s/%s", pkgName, v)
		full, err := template.ASTTemplates.Open(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		f, err := parser.ParseFile(token.NewFileSet(), "pkg/transform/templates/"+v, full, parser.DeclarationErrors)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outfset := token.NewFileSet()
		buf := &bytes.Buffer{}
		err = format.Node(buf, outfset, f)
		imps, err := imports.Process(fileName, buf.Bytes(), nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		imps = []byte(strings.ReplaceAll(string(imps), "foo", pkgName))
		imps = []byte(strings.ReplaceAll(string(imps), "Bar", "Index"))
		imps = []byte(strings.ReplaceAll(string(imps), "bar", "index"))

		if err = splatFile(fileName, bytes.NewBuffer(imps)); err != nil {
			fmt.Println(err)
		}
	}
}

func New() Service {
	return &service{}
}

func splatFile(target string, buf io.Reader) error {
	err := os.MkdirAll(path.Dir(target), os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "Couldn't create directory for %q", target)
	}
	f, err := os.Create(target)
	if err != nil {
		return errors.Wrapf(err, "Couldn't create file %q", target)
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = io.Copy(f, buf)
	return errors.Wrapf(err, "Error writing data to file %q", target)
}
