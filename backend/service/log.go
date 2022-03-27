package service

import (
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
	"go.uber.org/zap"
	"time"
)

func CreateLog(log *model.Log) (err error) {
	log.ID = nil

	cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	tx := cli.Table("log").Create(log)
	err = tx.Error
	if err == nil {
		xlog.Logger.Info("insert log", zap.Any("log", log))
	}
	return
}

func CreateBlob(log *model.Blob) (err error) {
	log.ID = nil

	cli, err := mysqlcli.GetMysqlClient()
	if err != nil {

		return
	}
	tx := cli.Table("log").Create(log)
	err = tx.Error
	if err == nil {
		xlog.Logger.Info("insert blob", zap.Any("log", log.Log), zap.Int("size", len(log.Blob)))
	}
	return
}
func GetBlob(id int) (b model.Blob, err error) {

	cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	tx := cli.Table("log").Where("id=?", id).First(&b)
	err = tx.Error

	return
}
func GetLogs(start, end time.Time, location, quadrant, atype, title, detail, review string, offset, limit int, orderby string) (bs []model.Log, err error) {
	cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	tx := cli.Table("log")
	if quadrant != "" {
		tx = tx.Where("quadrant like ?", quadrant)
	}
	if start.Unix() != (time.Time{}.Unix()) {
		tx = tx.Where("ctime >= ?", start)
	}
	if end.Unix() != (time.Time{}.Unix()) {
		tx = tx.Where("ctime <= ?", end)
	}
	if location != "" {
		//"name LIKE ?", "%jin%"
		tx = tx.Where("location like ?", "%"+location+"%")
	}
	if atype != "" {
		tx = tx.Where("atype like ?", "%"+atype+"%")
	}
	if title != "" {
		tx = tx.Where("title like ?", "%"+title+"%")
	}
	if detail != "" {
		tx = tx.Where("detail like ?", "%"+detail+"%")
	}
	if review != "" {
		tx = tx.Where("review like ?", "%"+review+"%")
	}
	tx = tx.Offset(offset).Limit(limit)
	if orderby != "" {
		tx = tx.Order(orderby)
	} else {
		tx = tx.Order("ctime desc")
	}
	tx = tx.Find(&bs)
	err = tx.Error
	return
}
