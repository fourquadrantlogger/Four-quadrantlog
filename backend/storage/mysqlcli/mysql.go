package mysqlcli

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func GetMysqlClient() (_ *gorm.DB, err error) {
	if _db == nil {
		_db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: "root:root@tcp(127.0.0.1:3306)/cplug?charset=utf8mb4&parseTime=True&loc=Local",
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}
		return _db, err
	} else {
		return _db, nil
	}
}
