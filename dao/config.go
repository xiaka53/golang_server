package dao

import "strconv"

const (
	AGENT_PRICE = "agent_price"
)

type Config struct {
	Id   int    `json:"id" gorm:"column:id;primary_key" description:""`
	Name string `json:"name" gorm:"column:id;primary_key" description:""`
	V    string `json:"v" gorm:"column:id;primary_key" description:""`
}

func (c *Config) TableName() string {
	return "config"
}

// 初始化配置 TODO
func (c *Config) init() {
	// mysqlConn.Table(c.TableName()).Where("`name`='?'", AGENT_PRICE).First(c)
}

// 获取代理商价格
func GetAgentPrice() float32 {
	var c *Config
	mysqlConn.Table(c.TableName()).Where("`name`='?'", AGENT_PRICE).First(c)
	if c.V == "" {
		return 0
	}
	price, _ := strconv.ParseFloat(c.V, 64)
	return float32(price)
}
