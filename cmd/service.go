/**
 * @Time : 2020/8/24 12:16 PM
 * @Author : solacowa@gmail.com
 * @File : service
 * @Software: GoLand
 */

package cmd

import (
	"context"
	"log"
	"path"

	"github.com/spf13/cobra"

	"github.com/icowan/kit-cli/pkg/service"
)

var serviceCmd = &cobra.Command{
	Use:     "service",
	Short:   "Generate new service",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("您必须提供服务名称")
			return
		}

		g := service.New(logger, args[0], path.Join("src", "pkg", "%s"), genSvc)
		if err := g.Gen(context.Background()); err != nil {
			log.Fatal("您必须提供服务名称")
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return prepare()
	},
}

func init() {
	newCmd.AddCommand(serviceCmd)
}
