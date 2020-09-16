/**
 * @Time : 2020/9/16 11:19 AM
 * @Author : solacowa@gmail.com
 * @File : service_test
 * @Software: GoLand
 */

package transform

import (
	"context"
	"testing"
)

var (
	svc = New()
)

func TestService_Init(t *testing.T) {
	ctx := context.Background()

	svc.Init(ctx, "world")

}

func TestService_TransformAST2(t *testing.T) {
	ctx := context.Background()
	svc.TransformAST(ctx, "hello", "http", "grpc")
}
