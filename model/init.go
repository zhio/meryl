package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"meryl/util"
	"os"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	//Error
	if connString == "" || err != nil {
		util.Log().Error("mysql host: %v", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("mysql lost: %v", err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	print("链接成功")
	DB = db
	//初始化
	migration()
}
