package mysqlcli

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var _db map[string]*gorm.DB = make(map[string]*gorm.DB)

func GetMysqlClient() (_ *gorm.DB, err error) {
	mysqlconn := os.Getenv("MYSQL")
	return MysqlClient(mysqlconn)

}
func MysqlClient(dsn string) (_ *gorm.DB, err error) {
	if _db[dsn] == nil {
		_db[dsn], err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}
		return _db[dsn], err
	} else {
		return _db[dsn], nil
	}
}
