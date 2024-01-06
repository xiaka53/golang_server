package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	mysqlConn *gorm.DB
	domain    string
)

func init() {
	dsn := os.Getenv("GOLANG_SERVER_DSN_STR")
	if dsn == "" {
		log.Fatalf("环境变量设置有误，无法启动服务DNS")
	}
	domain = os.Getenv("GOLANG_SERVER_DOMAIN_STR")
	if domain == "" {
		log.Fatalf("环境变量设置有误，无法启动服务DOMAIN")
	}
	dbPoolConfig := mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})

	gormConfig := gorm.Config{}
	db, err := gorm.Open(dbPoolConfig, &gormConfig, &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// 自动创建表格，如果表格已经存在，检查字段是否有变化
	if err := mysqlConn.AutoMigrate(
		&Account{},
	); err != nil {
		log.Fatalf("表格初始化失败：%v", err)
	}

	mysqlConn = db
}

// 获取域名
func GetDomain(url string) string {
	return url + domain
}
