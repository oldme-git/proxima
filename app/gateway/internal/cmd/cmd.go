package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"proxima/app/gateway/internal/controller/user"
	"proxima/app/gateway/internal/controller/words"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Bind(user.NewV1())
						group.Bind(words.NewV1())
					})
				})
			})
			s.Run()
			return nil
		},
	}
)