package models

import (
	"github.com/provider-go/pkg/mysql/typemysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewMysql(cfg typemysql.ConfigMysql) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level
			Colorful:                  true,        // 禁用彩色打印
			IgnoreRecordNotFoundError: true,        // 忽略记录不存在的错误
			ParameterizedQueries:      true,        // SQL日志中包含params
		},
	)
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns) // 设置最大空闲数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) // 设置最大连接数

	return db, nil
}
