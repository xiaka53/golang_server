package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var mysqlConn *gorm.DB

func init() {
	dsn := os.Getenv("GOLANG_SERVER_DSN_STR")
	if dsn == "" {
		log.Fatalf("环境变量设置有误，无法启动服务")
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
	mysqlConn = db
}
