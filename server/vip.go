package server

import (
	"context"
	"github.com/xiaka53/golang_server/dao"
	"github.com/xiaka53/golang_server/pkg/pkg"
)

type vip struct {
}

func NewVipServer() pkg.VipServer {
	return &vip{}
}

// 列表
func (v vip) List(ctx context.Context, req *pkg.VipListRepuest) (*pkg.VipListResponse, error) {
	var (
		info  pkg.VipListResponse
		data  dao.Account
		list  []dao.Vip
		state pkg.VipInfo_VipStatus
		err   error
	)
	data.Token = req.GetToken()
	if len(data.Token) != dao.TOKEN_LENGTH {
		goto OUT
	}
	if err = (&data).First(); err != nil {
		goto OUT
	}
	state = pkg.VipInfo_Unspecified
	if data.Id != dao.ADMIN_MARK {
		state = pkg.VipInfo_On
	}

	info.TotalCount, list = (&dao.Vip{Status: state, Name: req.GetName()}).GetVipList(req.From-1, req.Page)
	for _, v := range list {
		info.List = append(info.List, &pkg.VipInfo{
			Logo:          v.Logo,
			Name:          v.Name,
			Price:         v.Price,
			Msg:           v.Msg,
			Mark:          v.Id,
			Count:         v.OpenNum,
			Status:        v.Status,
			OriginalPrice: v.OriginalPrice,
		})
	}
OUT:
	return &info, err
}

// 执行开通
func (v vip) Open(ctx context.Context, req *pkg.OpenVipRepuest) (*pkg.VipResponse, error) {
	var (
		info  pkg.VipResponse
		data  dao.Account
		_vip  dao.Vip
		order dao.Order
		err   error
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
	_vip.Id = req.GetMark()
	if _vip.Id < 1 {
		info.Msg = "不存在的类目"
		goto OUT
	}
	if err = (&_vip).GetVipByMark(); err != nil {
		info.Msg = "获取类目失败"
		goto OUT
	}
	if data.Amount < _vip.Price {
		info.Msg = "您的预留资金不足，请充值"
		goto OUT
	}
	order = dao.Order{
		Aid:     data.Id,
		Agent:   data.Account,
		Account: req.GetAccount(),
		Vip:     _vip.Name,
		VipId:   _vip.Id,
		Price:   _vip.Price,
		Memo:    req.GetMemo(),
	}
	if err = (&order).CreateOrder(); err != nil {
		info.Msg = "开通失败"
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}

// 添加会员种类
func (v vip) AddVip(ctx context.Context, req *pkg.AddVipRequest) (*pkg.VipResponse, error) {
	var (
		info pkg.VipResponse
		data dao.Account
		_vip dao.Vip
		err  error
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
		info.Msg = "您无权执行此操作"
		goto OUT
	}
	_vip = dao.Vip{
		Name:          req.GetName(),
		Logo:          req.GetLogo(),
		OriginalPrice: req.GetOriginalPrice(),
		Price:         req.GetPrice(),
		Msg:           req.GetMsg(),
		Status:        pkg.VipInfo_On,
	}
	if msg := (&_vip).CreateVip(); msg != "" {
		info.Msg = msg
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}

// 开启会员
func (v vip) OnVip(ctx context.Context, req *pkg.OpenOrCloseVipRequest) (*pkg.VipResponse, error) {
	var (
		info pkg.VipResponse
		data dao.Account
		_vip dao.Vip
		err  error
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
	_vip.Id = req.GetMark()
	if _vip.Id < 1 {
		info.Msg = "不存在的类目"
		goto OUT
	}
	if err = (&_vip).GetVipByMark(); err != nil {
		info.Msg = "获取类目失败"
		goto OUT
	}
	if _vip.Status != pkg.VipInfo_On {
		info.Msg = "该类目已开启"
		goto OUT
	}
	if data.Id != dao.ADMIN_MARK {
		info.Msg = "您无权执行此操作"
		goto OUT
	}
	if err = (&_vip).OpenVip(); err != nil {
		info.Msg = "操作失败"
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}

// 修改会员信息
func (v vip) Edit(ctx context.Context, req *pkg.EditVipRequest) (*pkg.VipResponse, error) {
	var (
		info pkg.VipResponse
		data dao.Account
		_vip dao.Vip
		err  error
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
		info.Msg = "您无权执行此操作"
		goto OUT
	}
	_vip.Id = req.GetMark()
	if _vip.Id < 1 {
		info.Msg = "请修改正确的类目"
		goto OUT
	}
	if err = (&_vip).First(); err != nil {
		info.Msg = "不存在的类目"
		goto OUT
	}
	_vip.Name = req.GetName()
	_vip.Logo = req.GetLogo()
	_vip.OriginalPrice = req.GetOriginalPrice()
	_vip.Price = req.GetPrice()
	_vip.Msg = req.GetMsg()
	if msg := (&_vip).UpdateVip(); msg != "" {
		info.Msg = msg
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}

// 关闭会员
func (v vip) OffVip(ctx context.Context, req *pkg.OpenOrCloseVipRequest) (*pkg.VipResponse, error) {
	var (
		info pkg.VipResponse
		data dao.Account
		_vip dao.Vip
		err  error
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
	_vip.Id = req.GetMark()
	if _vip.Id < 1 {
		info.Msg = "不存在的类目"
		goto OUT
	}
	if err = (&_vip).GetVipByMark(); err != nil {
		info.Msg = "获取类目失败"
		goto OUT
	}
	if _vip.Status != pkg.VipInfo_Off {
		info.Msg = "该类目已关闭"
		goto OUT
	}
	if data.Id != dao.ADMIN_MARK {
		info.Msg = "您无权执行此操作"
		goto OUT
	}
	if err = (&_vip).CloseVip(); err != nil {
		info.Msg = "操作失败"
		goto OUT
	}
	info.Status = true
OUT:
	return &info, err
}
