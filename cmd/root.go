/**
 * @Time : 2020/8/24 12:06 PM
 * @Author : solacowa@gmail.com
 * @File : root
 * @Software: GoLand
 */

package cmd

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/icowan/kit-cli/pkg/fs"
	"github.com/icowan/kit-cli/pkg/generate"
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gokit",
		Short: "Go-Kit CLI",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	genSvc generate.Service
	fsSvc  fs.Service
	logger kitlog.Logger
)

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "If you want to se the debug logs.")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "Force overide existing files without asking.")
	rootCmd.PersistentFlags().StringP("folder", "b", "", "If you want to specify the base folder of the project.")
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func prepare() error {
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)

	fsSvc = fs.New(logger, "./", "./", false)
	genSvc = generate.New(logger, fsSvc)
	return nil
}
