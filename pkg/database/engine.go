package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"nes/config"
	"sync"
	"time"
)

type DBEngine struct {
	*gorm.DB
}

var (
	once     sync.Once
	dbEngine *DBEngine
)

func NewDbEngine(opts *config.Database) *DBEngine {

	once.Do(func() {
		db, err := gorm.Open(mysql.Open(opts.Source()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		if err != nil {
			panic(err)
		}
		sqlDB, err := db.DB()

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

		dbEngine = &DBEngine{db}
	})
	return dbEngine
}
