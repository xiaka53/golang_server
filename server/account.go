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
		price         float32
		memo          string
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
	price = dao.GetAgentPrice()
	if data.Amount < price {
		info.Msg = "您的预留资金不足，请充值"
		err = errors.New("amount low")
		goto OUT
	}
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
	memo = req.GetMemo()
	if memo == "" {
		info.Msg = "昵称无法设置为空"
		err = errors.New("memo no null")
		goto OUT
	}

	if msg := (&newUser).Create(memo); msg != "" {
		info.Msg = msg
		err = errors.New("create err")
		goto OUT
	}
OUT:
	return &info, err
}

// 关闭用户
func (a *account) CloseAccount(ctx context.Context, req *pkg.CloseAccountRequest) (*pkg.ExecAccountResponse, error) {
	var (
		info           pkg.ExecAccountResponse
		data, execUser dao.Account
		err            error
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
	if data.Id != dao.ADMIN_MARK {
		info.Msg = "您无权限此操作"
		err = errors.New("auth err")
		goto OUT
	}
	execUser.Account = req.GetAccount()
	if len(execUser.Account) == 0 {
		info.Msg = "不存在的此账号"
		err = errors.New("no user")
		goto OUT
	}
	if execUser.State == dao.Account_state_Off {
		info.Msg = "此账号已经关闭"
		err = errors.New("user is off")
		goto OUT
	}
	if err = execUser.Close(); err != nil {
		info.Msg = "关闭失败"
		goto OUT
	}
OUT:
	return &info, err
}

// 开启用户
func (a *account) OpenAccount(ctx context.Context, req *pkg.OpenAccountRequest) (*pkg.ExecAccountResponse, error) {
	var (
		info           pkg.ExecAccountResponse
		data, execUser dao.Account
		err            error
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
	if data.Id != dao.ADMIN_MARK {
		info.Msg = "您无权限此操作"
		err = errors.New("auth err")
		goto OUT
	}
	execUser.Account = req.GetAccount()
	if len(execUser.Account) == 0 {
		info.Msg = "不存在的此账号"
		err = errors.New("no user")
		goto OUT
	}
	if execUser.State == dao.Account_state_On {
		info.Msg = "此账号已经开启"
		err = errors.New("user is on")
		goto OUT
	}
	if err = execUser.Open(); err != nil {
		info.Msg = "开启失败"
		goto OUT
	}
OUT:
	return &info, err
}

// 个人资料查询
func (a *account) AccountInfo(ctx context.Context, req *pkg.AccountInfoRequest) (*pkg.AccountInfoResponse, error) {
	var (
		info pkg.AccountInfoResponse
		data dao.Account
		err  error
	)
	data.Token = req.GetToken()
	if len(data.Token) != dao.TOKEN_LENGTH {
		goto OUT
	}
	if err = (&data).First(); err != nil {
		goto OUT
	}
OUT:
	return &info, err
}

// 设置备注
func (a *account) SetMemoAccount(ctx context.Context, req *pkg.MemoRequest) (*pkg.ExecAccountResponse, error) {
	var (
		info           pkg.ExecAccountResponse
		data, execUser dao.Account
		memo           dao.Memo
		err            error
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
	execUser.Account = req.GetAccount()
	if len(execUser.Account) == 0 {
		info.Msg = "不存在的此账号"
		err = errors.New("no user")
		goto OUT
	}
	if execUser.Recommender != data.Id {
		info.Msg = "您无此权限"
		err = errors.New("no auth")
		goto OUT
	}
	memo.Aid = data.Id
	memo.Rid = execUser.Id
	memo.Memo = req.GetMemo()
	if memo.Memo == "" {
		info.Msg = "无法设置为空"
		err = errors.New("no null")
		goto OUT
	}
	if err := (&memo).Update(); err != nil {
		info.Msg = "修改失败"
		err = errors.New("update err")
		goto OUT
	}
OUT:
	return &info, err
}
