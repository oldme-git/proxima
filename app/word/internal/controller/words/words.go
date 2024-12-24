// Package words implements the gRPC controller layer for word management
package words

import (
	"context"

	v1 "proxima/app/word/api/words/v1"
	"proxima/app/word/internal/logic/words"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct {
	v1.UnimplementedWordsServer
	words *words.Words
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterWordsServer(s.Server, &Controller{
		words: words.New(),
	})
}

func (c *Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.words.Create(ctx, words.CreateInput{
		Uid:        req.Uid,
		Word:       req.Word,
		Definition: req.Definition,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{
		Id: uint32(id),
	}, nil
}

func (c *Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	word, err := c.words.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetRes{}
	err = gconv.Scan(word, &res.Words)
	return
}
