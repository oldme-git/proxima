package main

import (
    "github.com/gogf/gf/contrib/registry/etcd/v2"
    "github.com/gogf/gf/contrib/rpc/grpcx/v2"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/os/gctx"

    "proxima/app/gateway/internal/cmd"
)

func main() {
    var (
        ctx         = gctx.GetInitCtx()
        etcdAddress = g.Cfg("etcd").MustGet(ctx, "etcd.address").String()
    )
    grpcx.Resolver.Register(etcd.New(etcdAddress))

    cmd.Main.Run(ctx)
}
