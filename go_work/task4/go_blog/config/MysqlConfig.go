package config

import (
	"fmt"
	"github/metanode/go_blog/model/entity"

	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMysql() {

	mysqConf := GlobalConfig.Mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		mysqConf.UserName, mysqConf.Password, mysqConf.Host, mysqConf.Port, mysqConf.DBName, mysqConf.Timeout)

	shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用事务
		// 设置日志级别 全局生效
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().In(shanghaiLoc) // 指定时区生成时间
		},
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// 配置连接池
	sqlDB.SetMaxOpenConns(mysqConf.MaxOpenConns)
	sqlDB.SetMaxIdleConns(mysqConf.MaxIdleConns)
	duration, _ := time.ParseDuration(mysqConf.ConnMaxLifetime)
	sqlDB.SetConnMaxLifetime(duration)

	// 创建表
	if err = db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
	); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	DB = db
}
