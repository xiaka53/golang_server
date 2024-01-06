package dao

import "errors"

const (
	TOKEN_LENGTH = 20
	// 顶级账号
	ADMIN_MARK = 1
)

type Account struct {
	Id          int    `json:"id" gorm:"column:id;primary_key" description:""`
	Recommender int    `json:"recommender" gorm:"column:recommender" description:"推荐人"`
	Account     string `json:"account" gorm:"column:account" description:"账号"`
	Email       string `json:"email" gorm:"column:email" description:"邮箱"`
	Phone       string `json:"phone" gorm:"column:phone" description:"手机号"`
	Name        string `json:"name" gorm:"column:name" description:"名称"`
	Url         string `json:"url" gorm:"column:url" description:"URL"`
	Token       string `json:"token" gorm:"column:token" description:"token"`
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

// 创建新用户
func (a *Account) Create() string {
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

	if mysqlConn.Table(a.TableName()).Create(a).Error == nil {
		return ""
	}
	return "注册失败"
}
