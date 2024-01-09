package dao

import (
	"fmt"
	"github.com/xiaka53/golang_server/pkg/pkg"
	"time"
)

type Order struct {
	Id         int32     `json:"id" gorm:"column:id;primary_key" description:""`
	Aid        int       `json:"aid" gorm:"column:aid" description:"代理商ID"`
	Agent      string    `json:"agent" gorm:"column:agent" description:"代理商账号"`
	Account    string    `json:"account" gorm:"column:account" description:"开通账号"`
	Vip        string    `json:"vip" gorm:"column:vip" description:"开通vip类型"`
	VipId      int32     `json:"vip_id" gorm:"column:vip_id" description:"开通vipId"`
	Price      float32   `json:"price" gorm:"column:price" description:"价格"`
	State      pkg.State `json:"state" gorm:"column:state" description:"状态"`
	ExecCount  int32     `json:"exec_count" gorm:"column:exec_count" description:"执行次数"`
	Msg        string    `json:"msg" gorm:"column:msg" description:"申诉内容"`
	Memo       string    `json:"memo" gorm:"column:memo" description:"备注"`
	CreateDate int64     `json:"create_date" gorm:"column:create_date" description:"执行时间"`
	UpdateDate int64     `json:"update_date" gorm:"column:create_date" description:"修改时间"`
}

func (o *Order) TableName() string {
	return "order"
}

// 根据mark获取订单
func (o *Order) GetOrderByMark() error {
	return mysqlConn.Table(o.TableName()).Where("id=?", o.Id).First(o).Error
}

// 获取订单分页
func (o *Order) GetOrder(from, amount int32) (total int32, data []Order) {
	var (
		_total     int64
		whereState string = "1=1"
		whereAgent string = "1=1"
	)
	if o.Aid > 0 {
		whereAgent = fmt.Sprintf("agent='%v'", o.Agent)
	}
	if o.State != pkg.State_Unspecified {
		whereState = fmt.Sprintf("`state`=%v", o.State)
	}
	if mysqlConn.Table(o.TableName()).Where(whereAgent).Where(whereState).Limit(int(amount)).Order("id desc").Offset(int(amount*from)).Find(&data).Error != nil {
		return 0, nil
	}
	if mysqlConn.Table(o.TableName()).Where(whereAgent).Where(whereState).Count(&_total).Error != nil {
		return 0, nil
	}
	total = int32(_total)
	return
}

// 创建订单
func (o *Order) CreateOrder() error {
	o.State = pkg.State_Review
	o.CreateDate = time.Now().Unix()
	db := mysqlConn.Begin()
	if err := db.Table(o.TableName()).Create(o).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := (&Account{Id: o.Aid}).funds(db, -o.Price, Detail_Type_BuyVip); err != nil {
		db.Rollback()
		return err
	}
	if err := (&Vip{Id: o.VipId}).addBuyVip(db); err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// 取消订单
func (o *Order) CancelOrder() error {
	db := mysqlConn.Begin()
	if err := db.Table(o.TableName()).Where("id=? and `state` in (2,5)", o.Id).Updates(map[string]any{
		"`state`":     pkg.State_Cancel,
		"update_date": time.Now().Unix(),
	}).Error; err != nil {
		db.Rollback()
		return err
	}
	// 余额返还
	if err := (&Account{Id: o.Aid}).funds(db, o.Price, Detail_Type_CancelOrder); err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// 完成订单
func (o *Order) FinishOrder() error {
	db := mysqlConn.Begin()
	if err := db.Table(o.TableName()).Where("id=? and `state` in (2,5)", o.Id).Updates(map[string]any{
		"`state`":     pkg.State_Finish,
		"update_date": time.Now().Unix(),
	}).Error; err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// 申请售后
func (o *Order) ApplyAfter() error {
	db := mysqlConn.Begin()
	if err := db.Table(o.TableName()).Where("id=? and `state` in (1,4)", o.Id).Updates(map[string]any{
		"`state`":     pkg.State_Appeal,
		"msg":         o.Msg,
		"update_date": time.Now().Unix(),
	}).Error; err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// 修改备注
func (o *Order) UpdateMemo() error {
	db := mysqlConn.Begin()
	if err := db.Table(o.TableName()).Where("id=?", o.Id).Updates(map[string]any{
		"memo":        o.Memo,
		"update_date": time.Now().Unix(),
	}).Error; err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
