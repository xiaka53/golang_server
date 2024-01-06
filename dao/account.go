package dao

type Account struct {
	Id          int    `json:"id" gorm:"cloumn:id;primary_key" description:""`
	Recommender int    `json:"recommender" gorm:"cloumn:recommender" description:"推荐人"`
	Account     string `json:"account" gorm:"cloumn:account" description:"账号"`
	Email       string `json:"email" gorm:"cloumn:email" description:"邮箱"`
	Phone       string `json:"phone" gorm:"cloumn:phone" description:"手机号"`
	Name        string `json:"name" gorm:"cloumn:name" description:"名称"`
	Url         string `json:"url" gorm:"cloumn:url" description:"URL"`
}

func (a *Account) TableName() string {
	return "account"
}
