package server

import (
	"context"
	"errors"
	"github.com/xiaka53/golang_server/dao"
	"github.com/xiaka53/golang_server/pkg/pkg"
)

type money struct {
}

func NewMoneyServer() pkg.MoneyServer {
	return &money{}
}

// 执行充值
func (m *money) Rechange(ctx context.Context, req *pkg.RechangeRepuest) (*pkg.ExecMoneyResponse, error) {
	var (
		info               pkg.ExecMoneyResponse
		data, rechangeData dao.Account
		rechangeNum        float32
		err                error
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
		info.Msg = "您无权限给代理商充值"
		err = errors.New("rechange err")
		goto OUT
	}
	rechangeNum = req.GetAmount()
	if rechangeNum < 1 {
		info.Msg = "最低充值1元"
		err = errors.New("rechange min 1RNB")
		goto OUT
	}
	if data.Amount < req.GetAmount() {
		info.Msg = "您的预留资金不足，请充值"
		err = errors.New("amount low")
		goto OUT
	}
	rechangeData.Account = req.GetAccount()
	if rechangeData.Account == "" {
		info.Msg = "不存在的账号"
		err = errors.New("no user account")
		goto OUT
	}
	_ = (&rechangeData).First()
	if rechangeData.Id < 1 {
		info.Msg = "不存在的用户"
		err = errors.New("no user")
		goto OUT
	}
	if rechangeData.Recommender != data.Id {
		info.Msg = "您无权限给此代理商充值"
		err = errors.New("rechange err this agent")
		goto OUT
	}
	if msg := (&data).Rechange(rechangeData.Id, rechangeNum); msg != "" {
		info.Msg = msg
		err = errors.New("rechange err")
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}

// 申请退款
func (m *money) RefundMoney(ctx context.Context, req *pkg.RefundMoneyRequest) (*pkg.ExecMoneyResponse, error) {
	var (
		info             pkg.ExecMoneyResponse
		data, refundData dao.Account
		refundNum        float32
		err              error
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
		info.Msg = "您无权限给代理商退款"
		err = errors.New("refund err")
		goto OUT
	}
	refundNum = req.GetAmount()
	if refundNum < 1 {
		info.Msg = "最低退款1元"
		err = errors.New("refund min 1RNB")
		goto OUT
	}
	refundData.Account = req.GetAccount()
	if refundData.Account == "" {
		info.Msg = "不存在的账号"
		err = errors.New("no user account")
		goto OUT
	}
	_ = (&refundData).First()
	if refundData.Id < 1 {
		info.Msg = "不存在的用户"
		err = errors.New("no user")
		goto OUT
	}
	if refundData.Recommender != data.Id {
		info.Msg = "您无权限给此代理商退款"
		err = errors.New("refund err this agent")
		goto OUT
	}
	if refundData.Amount < req.GetAmount() {
		info.Msg = "该代理商的预留资金不足"
		err = errors.New("amount low")
		goto OUT
	}
	if msg := (&data).Rechange(refundData.Id, refundNum); msg != "" {
		info.Msg = msg
		err = errors.New("refund err")
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}
