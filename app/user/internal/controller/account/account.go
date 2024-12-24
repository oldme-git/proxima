// Package account implements the gRPC controller layer for user account management
package account

import (
    "context"

    v1 "proxima/app/user/api/account/v1"
    "proxima/app/user/internal/logic/account"

    "github.com/gogf/gf/contrib/rpc/grpcx/v2"
    "github.com/gogf/gf/v2/util/gconv"
)

// Controller implements the gRPC server interface for account management
// It acts as a bridge between the gRPC transport layer and business logic layer
type Controller struct {
    v1.UnimplementedAccountServer
    account *account.Account // Business logic layer instance
}

// Register registers the account service with the gRPC server
// This function is called during server initialization
func Register(s *grpcx.GrpcServer) {
    v1.RegisterAccountServer(s.Server, &Controller{
        account: account.New(),
    })
}

// Register handles user registration requests
// It converts the gRPC request to business logic input and returns the new user's ID
func (c *Controller) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
    id, err := c.account.Register(ctx, account.RegisterInput{
        Passport: req.Passport,
        Password: req.Password,
        Email:    req.Email,
    })
    if err != nil {
        return nil, err
    }
    return &v1.RegisterRes{
        Id: int32(id),
    }, nil
}

// Login handles user authentication requests
// It validates credentials and returns a JWT token upon successful authentication
func (c *Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
    token, err := c.account.Login(ctx, req.Passport, req.Passport)
    if err != nil {
        return nil, err
    }
    return &v1.LoginRes{
        Token: token,
    }, nil
}

// Info retrieves user information using a JWT token
// It validates the token and returns the user's profile information
func (c *Controller) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
    entityData, err := c.account.Info(ctx, req.Token)
    if err != nil {
        return nil, err
    }
    res = &v1.InfoRes{}
    err = gconv.Scan(entityData, &res.User)
    return
}
