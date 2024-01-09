package dao

import (
	"fmt"
	"github.com/xiaka53/golang_server/pkg/pkg"
	"gorm.io/gorm"
)

type Vip struct {
	Id            int32                 `json:"id" gorm:"column:id;primary_key" description:""`
	Name          string                `json:"name" gorm:"column:name" description:"名称"`
	Logo          string                `json:"logo" gorm:"column:logo" description:"logo"`
	OriginalPrice float32               `json:"original_price" gorm:"column:original_price" description:"原价"`
	Price         float32               `json:"price" gorm:"column:price" description:"价格"`
	Msg           string                `json:"msg" gorm:"column:msg" description:"备注"`
	OpenNum       int32                 `json:"open_num" gorm:"column:open_num" description:"开通数量"`
	Status        pkg.VipInfo_VipStatus `json:"status" gorm:"column:status" description:"状态"`
}

func (v *Vip) TableName() string {
	return "vip"
}

func (v *Vip) First() error {
	return mysqlConn.Table(v.TableName()).Where(v).First(v).Error
}

// 根据mark获取VIP
func (v *Vip) GetVipByMark() error {
	return mysqlConn.Table(v.TableName()).Where("id=?", v.Id).First(v).Error
}

// 获取Vip列表
func (v *Vip) GetVipList(from, amount int32) (total int32, data []Vip) {
	var (
		_total      int64
		whereStatus string = "1=1"
		whereName   string = "1=1"
	)
	if v.Status != pkg.VipInfo_Unspecified {
		whereStatus = fmt.Sprintf("`status`= %v", v.Status)
	}
	if v.Name != "" {
		whereName = fmt.Sprintf("`name` like `%%%v%%`", v.Name)
	}
	if mysqlConn.Table(v.TableName()).Where(whereStatus).Where(whereName).Limit(int(amount)).Order("id desc").Offset(int(amount*from)).Find(&data).Error != nil {
		return 0, nil
	}
	if mysqlConn.Table(v.TableName()).Where(whereStatus).Where(whereName).Count(&_total).Error != nil {
		return 0, nil
	}
	total = int32(_total)
	return
}

// 创建会员
func (v *Vip) CreateVip() string {
	var (
		_nameVip = Vip{Name: v.Name}
	)
	_ = (&_nameVip).First()
	if _nameVip.Id > 0 {
		return "存在的名称"
	}
	if mysqlConn.Table(v.TableName()).Create(v).Error != nil {
		return "创建失败"
	}
	return ""
}

// 修改会员
func (v *Vip) UpdateVip() string {
	var (
		_nameVip = Vip{Name: v.Name}
	)
	_ = (&_nameVip).First()
	if _nameVip.Id > 0 && _nameVip.Id != v.Id {
		return "存在的名称"
	}
	if mysqlConn.Table(v.TableName()).Updates(v).Error != nil {
		return "创建失败"
	}
	return ""
}

// 开启会员
func (v *Vip) OpenVip() error {
	return mysqlConn.Table(v.TableName()).Where("`status`=? and id=?", pkg.VipInfo_Off, v.Id).Update("`status`", pkg.VipInfo_On).Error
}

// 关闭会员
func (v *Vip) CloseVip() error {
	return mysqlConn.Table(v.TableName()).Where("`status`=? and id=?", pkg.VipInfo_On, v.Id).Update("`status`", pkg.VipInfo_Off).Error
}

// 会员增加购买
func (v *Vip) addBuyVip(db *gorm.DB) error {
	return db.Table(v.TableName()).Where("id=?", v.Id).Update("`open_num`", gorm.Expr("open_num+1")).Error
}

// 会员减少购买
func (v *Vip) subBuyVip(db *gorm.DB) error {
	return db.Table(v.TableName()).Where("id=?", v.Id).Update("`open_num`", gorm.Expr("open_num-1")).Error
}
