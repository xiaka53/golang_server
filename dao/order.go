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

// 取消订单
func (o *Order) CancelOrder() error {
	db := mysqlConn.Begin()
	if err := db.Table(o.TableName()).Where("id=? and `state` in (2,5)").Updates(map[string]interface{}{
		"`state`":     pkg.State_Cancel,
		"update_date": time.Now().Unix(),
	}); err != nil {
		db.Rollback()
	}
	// 余额返还
	db.Commit()
	return nil
}
