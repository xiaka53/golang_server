package server

import (
	"context"
	"github.com/xiaka53/golang_server/dao"
	"github.com/xiaka53/golang_server/pkg/pkg"
)

type record struct {
}

func NewRecordServer() pkg.RecordServer {
	return &record{}
}

// 查询流水
func (r *record) List(ctx context.Context, req *pkg.RecordListRequest) (*pkg.RecordListResponse, error) {
	var (
		info  pkg.RecordListResponse
		data  dao.Account
		list  []dao.Order
		agent string
		err   error
	)
	data.Token = req.GetToken()
	if len(data.Token) != dao.TOKEN_LENGTH {
		goto OUT
	}
	if err = (&data).First(); err != nil {
		goto OUT
	}
	agent = req.GetAgent()
	if data.Id != dao.ADMIN_MARK && agent != "" {
		goto OUT
	}
	if agent == "" {
		agent = data.Account
	}
	info.TotalCount, list = (&dao.Order{Agent: agent, State: req.GetState()}).GetOrder(req.From-1, req.Page)
	for _, v := range list {
		info.List = append(info.List, &pkg.RecordInfo{
			Agent:     v.Agent,
			Account:   v.Account,
			Vip:       v.Vip,
			Price:     v.Price,
			State:     v.State,
			ExecCount: v.ExecCount,
			Msg:       v.Msg,
			Memo:      v.Memo,
			Mark:      v.Id,
		})
	}
OUT:
	return &info, err
}

// 取消充值
func (r *record) Cancel(ctx context.Context, req *pkg.CancelRecordRequest) (*pkg.RecordResponse, error) {
	var (
		info  pkg.RecordResponse
		data  dao.Account
		order dao.Order
		agent string
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
	order.Id = req.GetMark()
	if order.Id < 1 {
		info.Msg = "不存在的订单"
		goto OUT
	}
	if err = (&order).GetOrderByMark(); err != nil {
		info.Msg = "获取订单失败"
		goto OUT
	}
	if order.State != pkg.State_Review && order.State != pkg.State_Appeal {
		info.Msg = "该订单无法执行此操作"
		goto OUT
	}
	if data.Id != dao.ADMIN_MARK && data.Id != order.Aid {
		info.Msg = "您无权执行此订单"
		goto OUT
	}

OUT:
	return &info, err
}

// 充值完成
func (r *record) Finish(ctx context.Context, req *pkg.FinishRecordRequest) (*pkg.RecordResponse, error) {

}

// 申请售后
func (r *record) ApplyAfterSales(ctx context.Context, req *pkg.ApplyAfterSalesRequest) (*pkg.RecordResponse, error) {

}

// 填写备注
func (r *record) Memo(ctx context.Context, req *pkg.MemoRepuest) (*pkg.RecordResponse, error) {

}
