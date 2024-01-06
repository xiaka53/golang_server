package server

import (
	"context"
	"errors"
	"github.com/xiaka53/golang_server/dao"
	"github.com/xiaka53/golang_server/pkg/pkg"
	"github.com/xiaka53/golang_server/pkg/tool"
	"strings"
)

type account struct {
}

func NewAccountServer() pkg.AccountServer {
	return &account{}
}

// 账号查询是否匹配
func (a *account) AccountSelect(ctx context.Context, req *pkg.AccountSelectRequest) (*pkg.AccountSelectResponse, error) {
	var (
		info pkg.AccountSelectResponse
		data dao.Account
		err  error
	)
	info.Token = ""
	if req.Account == "" || req.Url == "" {
		goto OUT
	}
	data.Account = req.GetAccount()
	if err = (&data).First(); err != nil {
		goto OUT
	}
	if strings.ToLower(req.GetUrl()) != data.Url {
		goto OUT
	}
SET_TOKEN:
	if info.Token, err = tool.GenerateToken(data.Id, dao.TOKEN_LENGTH); err != nil {
		goto OUT
	}
	if data.Token == info.Token {
		goto SET_TOKEN
	}
	data.Token = info.Token
	if err = data.SetToken(); err != nil {
		goto OUT
	}
OUT:
	return &info, err
}

// 创建新用户
func (a *account) AddAccount(ctx context.Context, req *pkg.AddAccountRequest) (*pkg.ExecAccountResponse, error) {
	var (
		info          pkg.ExecAccountResponse
		data, newUser dao.Account
		err           error
	)
	data.Token = req.GetToken()
	if len(data.Token) != dao.TOKEN_LENGTH {
		info.Msg = "登陆错误"
		goto OUT
	}
	if err = (&data).First(); err != nil {
		info.Msg = "登陆超时"
		goto OUT
	}
	if data.Recommender != dao.ADMIN_MARK && data.Id != dao.ADMIN_MARK {
		info.Msg = "您无权限推荐代理商"
		err = errors.New("recommender err")
		goto OUT
	}
	// 资金判断 TODO

	newUser = dao.Account{
		Recommender: data.Id,
		Account:     req.GetAccount(),
		Email:       req.GetEmail(),
		Phone:       req.GetPhone(),
		Name:        req.GetName(),
		Url:         strings.ToLower(req.GetUrl()),
	}
	if tool.PingUrl(dao.GetDomain(newUser.Url)) == false {
		info.Msg = "请更换域名，无法使用此域名"
		err = errors.New("ping url err")
		goto OUT
	}
	if msg := (&newUser).Create(); msg != "" {
		info.Msg = msg
		err = errors.New("create err")
		goto OUT
	}
OUT:
	return &info, err
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
