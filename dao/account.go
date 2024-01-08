package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

const (
	TOKEN_LENGTH = 20
	// 顶级账号
	ADMIN_MARK = 1
	// 账号开启状态
	Account_state_On = "on"
	// 账号关闭状态
	Account_state_Off = "off"
)

type Account struct {
	Id          int     `json:"id" gorm:"column:id;primary_key" description:""`
	Recommender int     `json:"recommender" gorm:"column:recommender" description:"推荐人"`
	Account     string  `json:"account" gorm:"column:account" description:"账号"`
	Email       string  `json:"email" gorm:"column:email" description:"邮箱"`
	Phone       string  `json:"phone" gorm:"column:phone" description:"手机号"`
	Name        string  `json:"name" gorm:"column:name" description:"名称"`
	Url         string  `json:"url" gorm:"column:url" description:"URL"`
	State       string  `json:"state" gorm:"column:state" description:"状态"`
	Token       string  `json:"token" gorm:"column:token" description:"token"`
	Amount      float32 `json:"amount" gorm:"column:token" description:"token"`
}

func (a *Account) TableName() string {
	return "account"
}

// 获取一条数据
func (a *Account) First() error {
	return mysqlConn.Table(a.TableName()).Where(a).First(a).Error
}

// 设置token
func (a *Account) SetToken() error {
	var detection Account
	detection.Token = a.Token
	if err := (&detection).GetFirstByToken(); err != nil {
		return err
	}
	if detection.Id > 0 && detection.Id != a.Id {
		return errors.New("existing token")
	}
	return mysqlConn.Table(a.Token).Where("id=?", a.Id).Update("token", a.Token).Error
}

// 根据token获取用户
func (a *Account) GetFirstByToken() error {
	if a.Token == "" {
		return errors.New("no token")
	}
	if len(a.Token) != TOKEN_LENGTH {
		return errors.New("token length err")
	}
	return mysqlConn.Table(a.TableName()).Where("token='?'", a.Token).First(a).Error
}

// 根据推荐人获取用户
func (a *Account) GetRecommenderBuyAccount(from, amount int32) (total int32, data []Account) {
	var _total int64
	if mysqlConn.Table(a.TableName()).Where("recommender=?", a.Recommender).Limit(int(amount)).Order("id desc").Offset(int(amount*from)).Find(&data).Error != nil {
		return 0, nil
	}
	if mysqlConn.Table(a.TableName()).Where("recommender=?", a.Recommender).Count(&_total).Error != nil {
		return 0, nil
	}
	total = int32(_total)
	return
}

// 创建新用户
func (a *Account) Create(memo string) string {
	var detection Account
	detection.Account = a.Account
	_ = (&detection).First()
	if detection.Id > 0 {
		return "该账户名称已存在"
	}

	detection = Account{Email: a.Email}
	_ = (&detection).First()
	if detection.Id > 0 {
		return "该邮箱已绑定，请换一个"
	}

	detection = Account{Phone: a.Phone}
	_ = (&detection).First()
	if detection.Id > 0 {
		return "该手机号已绑定，请换一个"
	}

	detection = Account{Url: a.Url}
	_ = (&detection).First()
	if detection.Id > 0 {
		return "该域名已绑定，请换一个"
	}

	db := mysqlConn.Begin()
	if db.Table(a.TableName()).Create(a).Error != nil {
		db.Rollback()
		return "注册失败"
	}
	if err := (&Account{Id: a.Recommender}).funds(db, -GetAgentPrice(), Detail_Type_Recommender); err != nil {
		db.Rollback()
		return "资金扣除失败，请检查余额"
	}
	if err := (&Memo{Aid: a.Id, Rid: a.Recommender, Memo: memo}).create(db); err != nil {
		db.Rollback()
		return "注册失败了"
	}
	db.Commit()
	return ""
}

// 关闭用户
func (a *Account) Close() error {
	a.State = Account_state_On
	return mysqlConn.Table(a.TableName()).Where(a).Update("state", Account_state_Off).Error
}

// 开启用户
func (a *Account) Open() error {
	a.State = Account_state_Off
	return mysqlConn.Table(a.TableName()).Where(a).Update("state", Account_state_On).Error
}

// 充值
func (a *Account) Rechange(rechangeId int, num float32) string {
	db := mysqlConn.Begin()
	if err := (&Account{Id: rechangeId}).funds(db, num, Detail_Type_Rechange); err != nil {
		db.Rollback()
		return "资金充值失败，请检查余额"
	}
	if err := (&Account{Id: a.Id}).funds(db, -num, Detail_Type_RechangeDeduct); err != nil {
		db.Rollback()
		return "资金扣除失败，请检查余额"
	}
	db.Commit()
	return ""
}

// 退款
func (a *Account) Refund(rechangeId int, num float32) string {
	db := mysqlConn.Begin()
	if err := (&Account{Id: rechangeId}).funds(db, -num, Detail_Type_Withdrawal); err != nil {
		db.Rollback()
		return "资金扣除失败，请检查余额"
	}
	if err := (&Account{Id: a.Id}).funds(db, num, Detail_Type_WithdrawalAdd); err != nil {
		db.Rollback()
		return "资金回流失败，请检查余额"
	}
	db.Commit()
	return ""
}

// 资金操作
func (a *Account) funds(db *gorm.DB, amount float32, detailType int) error {
	if err := db.Exec(fmt.Sprintf("Update %v set amount+%v where id=%v", a.TableName(), amount, a.Id)).Error; err != nil {
		return err
	}
	if err := (&Detail{Aid: a.Id, Num: amount, Type: detailType}).create(db); err != nil {
		return err
	}
	return nil
}
