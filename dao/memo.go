package dao

import "gorm.io/gorm"

type Memo struct {
	Id   int    `json:"id" gorm:"column:id;primary_key" description:""`
	Aid  int    `json:"aid" gorm:"column:aid" description:"用户ID"`
	Rid  int    `json:"rid" gorm:"column:rid" description:"下级Id"`
	Memo string `json:"memo" gorm:"column:memo" description:"备注"`
}

func (m *Memo) TableName() string {
	return "memo"
}

func (m *Memo) create(db *gorm.DB) error {
	return db.Table(m.TableName()).Create(m).Error
}

// 修改昵称
func (m *Memo) Update() error {
	return mysqlConn.Table(m.TableName()).Where("aid=? and rid=?", m.Aid, m.Rid).Update("memo", m.Memo).Error
}

// 获取昵称
func (m *Memo) GetMemo() string {
	mysqlConn.Table(m.TableName()).Where("aid=? and rid=?", m.Aid, m.Rid).First(m)
	return m.Memo
}
