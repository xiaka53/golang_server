package server

import (
	"context"
	"github.com/xiaka53/golang_server/pkg/pkg"
)

type account struct {
}

func NewAccountServer() pkg.AccountServer {
	return &account{}
}

// 账号查询是否匹配
func (a *account) AccountSelect(ctx context.Context, req *pkg.AccountSelectRequest) (*pkg.AccountSelectResponse, error) {
	return nil, nil
}

// 创建新用户
func (a *account) AddAccount(ctx context.Context, req *pkg.AddAccountRequest) (*pkg.ExecAccountResponse, error) {
	return nil, nil
}

// 关闭用户
func (a *account) CloseAccount(ctx context.Context, req *pkg.CloseAccountRequest) (*pkg.ExecAccountResponse, error) {
	return nil, nil
}

// 开启新用户
func (a *account) OpenAccount(ctx context.Context, req *pkg.OpenAccountRequest) (*pkg.ExecAccountResponse, error) {
	return nil, nil
}

// 个人资料查询
func (a *account) AccountInfo(ctx context.Context, req *pkg.AccountInfoRequest) (*pkg.AccountInfoResponse, error) {
	return nil, nil
}
