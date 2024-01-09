package dao

import (
	"gorm.io/gorm"
	"time"
)

const (
	// 开户
	Detail_Type_Recommender = iota + 1
	// 购买
	Detail_Type_BuyVip
	// 提现
	Detail_Type_Withdrawal
	// 提现增加
	Detail_Type_WithdrawalAdd
	// 充值
	Detail_Type_Rechange
	// 充值扣除
	Detail_Type_RechangeDeduct
	// 取消订单
	Detail_Type_CancelOrder
)

type Detail struct {
	Id         int     `json:"id" gorm:"column:id;primary_key" description:""`
	Aid        int     `json:"aid" gorm:"column:id" description:"用户ID"`
	Amount     float32 `json:"amount" gorm:"column:id" description:"操作前金额"`
	Num        float32 `json:"num" gorm:"column:id" description:"操作金额"`
	Type       int     `json:"type" gorm:"column:type" description:"类别"`
	Remark     string  `json:"remark" gorm:"column:remark" description:"备注"`
	CreateDate int64   `json:"create_date" gorm:"column:create_date" description:"执行时间"`
}

func (d *Detail) TableName() string {
	return "detail"
}

// 创建流水
func (d *Detail) create(db *gorm.DB) error {
	var _account Account
	_account.Id = d.Aid
	_ = (&_account).First()
	d.Amount = _account.Amount
	d.CreateDate = time.Now().Unix()
	switch d.Type {
	case Detail_Type_Recommender: // 开户
	case Detail_Type_BuyVip: // 购买
	case Detail_Type_Withdrawal: // 提现
	case Detail_Type_WithdrawalAdd: // 提现增加
	case Detail_Type_Rechange: // 充值
	case Detail_Type_RechangeDeduct: // 充值扣除

	}
	return db.Table(d.TableName()).Create(d).Error
}
