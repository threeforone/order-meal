package initialize

import (
	"github.com/threeforone/order-meal/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Mysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(global.Conf.Mysql.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil
	}
	sql, _ := db.DB()
	sql.SetMaxOpenConns(global.Conf.Mysql.MaxOpenConns)
	sql.SetMaxIdleConns(global.Conf.Mysql.MaxIdleConns)
	return db
}
