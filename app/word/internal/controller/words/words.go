// Package words implements the gRPC controller layer for word management
package words

import (
    "context"

    v1 "proxima/app/word/api/words/v1"
    "proxima/app/word/internal/logic/words"

    "github.com/gogf/gf/contrib/rpc/grpcx/v2"
    "github.com/gogf/gf/v2/util/gconv"
)

// Controller implements the gRPC server interface for word management
// It acts as a bridge between the gRPC transport layer and business logic layer
type Controller struct {
    v1.UnimplementedWordsServer
    words *words.Words // Business logic layer instance
}

// Register registers the words service with the gRPC server
// This function is called during server initialization
func Register(s *grpcx.GrpcServer) {
    v1.RegisterWordsServer(s.Server, &Controller{
        words: words.New(),
    })
}

// Create handles word creation requests
// It converts the gRPC request to business logic input and returns the new word's ID
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

// Get retrieves word information by ID
// It returns the complete word entity including all its attributes
func (c *Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
    word, err := c.words.Get(ctx, req.Id)
    if err != nil {
        return nil, err
    }
    res = &v1.GetRes{}
    err = gconv.Scan(word, &res.Words)
    return
}
