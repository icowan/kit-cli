/**
 * @Time : 2020/8/24 12:10 PM
 * @Author : solacowa@gmail.com
 * @File : main
 * @Software: GoLand
 */

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/icowan/kit-cli/cmd"
	"github.com/icowan/kit-cli/util"
)

func main() {
	gosrc := strings.TrimSuffix(util.GetGOPATH(), afero.FilePathSeparator) + afero.FilePathSeparator + "src" + afero.FilePathSeparator
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	gosrc, err = filepath.EvalSymlinks(gosrc)
	if err != nil {
		log.Fatal(err)
		return
	}
	pwd, err = filepath.EvalSymlinks(pwd)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !strings.HasPrefix(pwd, gosrc) {
		log.Fatal("The project must be in the $GOPATH/src folder for the generator to work.")
		return
	}
	cmd.Run()
}
