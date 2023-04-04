package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"noir.github.com/aws_billing/internal/modules/logger"
	"noir.github.com/aws_billing/internal/modules/setting"
	"strings"
	"time"
)

var DB *gorm.DB

type Status int8
type CommonMap map[string]interface{}

const (
	Disabled Status = 0 // 禁用
	Failure  Status = 0 // 失败
	Enabled  Status = 1 // 启用
	Finish   Status = 2 // 完成
)

const DefaultLongTimeFormat = "2006-01-02 15:04:05"
const DefaultShortTimeFormat = "2006-01-02"

const (
	dbPingInterval = 90 * time.Second
	dbMaxLiftTime  = 2 * time.Hour
)

func CreateDB() *gorm.DB {
	s := setting.Read()
	dsn := getDbDSN(s)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(fmt.Sprintf("数据库连接失败-%s", err.Error()))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(s.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(s.DB.MaxOpenConns)

	// 创建表
	err = db.AutoMigrate(&User{}, &LoginLog{}, &Platform{})
	if err != nil {
		logger.Infof(fmt.Sprintf("数据库创建表失败-%s", err))
	}
	return db
}

func getDbDSN(setting *setting.Setting) string {
	engine := strings.ToLower(setting.DB.Engine)
	dsn := ""
	switch engine {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			setting.DB.User,
			setting.DB.Password,
			setting.DB.Host,
			setting.DB.Port,
			setting.DB.Database,
			setting.DB.Charset)
	}
	return dsn
}

func keepDBALive(db *gorm.DB) {
	t := time.Tick(dbPingInterval)
	var err error
	for {
		<-t
		sqlDB, _ := db.DB()
		err = sqlDB.Ping()
		if err != nil {
			logger.Infof("database ping: %s", err)
		}
	}
}
