package service

import (
	"errors"
	"fmt"
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

func DeleteLog(id string) (err error) {

	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	if id == "" {
		return errors.New("id not set")
	}
	tx := cli.Table("log").Where("id = ?", id).Delete(new(model.Log))
	err = tx.Error
	if err == nil {
		xlog.Logger.Info("insert log", zap.Any("log", id))
	}
	return
}
func UpdateLog(log *model.Log) (err error) {
	cfg, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	logmodel := Model(cfg.DBName, "log")
	if log.Detail != nil {
		detaillen, e := CheckStringLen(*log.Detail)
		if e != nil {
			return e
		}
		if detaillen.Len > logmodel["detail"].Length {
			return errors.New("detail text too long," + fmt.Sprint(detaillen.Len) + "/" + fmt.Sprint(logmodel["detail"].Length))
		}
	}
	tx := cli.Table("log").Where("id = ?", log.ID).Updates(log)
	err = tx.Error
	if err == nil {
		xlog.Logger.Info("insert log", zap.Any("log", log))
	}
	return
}
func CreateLog(log *model.Log) (err error) {
	log.ID = uuid.New().String()
	cfg, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	logmodel := Model(cfg.DBName, "log")
	if log.Detail != nil {
		detaillen, e := CheckStringLen(*log.Detail)
		if e != nil {
			return e
		}
		if detaillen.Len > logmodel["detail"].Length {
			return errors.New("detail text too long," + fmt.Sprint(detaillen.Len) + "/" + fmt.Sprint(logmodel["detail"].Length))
		}
	}
	if log.Review != nil {
		detaillen, _ := CheckStringLen(*log.Review)
		if detaillen.Len > logmodel["review"].Length {
			return errors.New("review text too long," + fmt.Sprint(detaillen.Len) + "/" + fmt.Sprint(logmodel["review"].Length))
		}
	}
	{
		detaillen, _ := CheckStringLen(log.Title)
		if detaillen.Len > logmodel["title"].Length {
			return errors.New("title text too long," + fmt.Sprint(detaillen.Len) + "/" + fmt.Sprint(logmodel["title"].Length))
		}
	}
	{
		detaillen, _ := CheckStringLen(*log.Location)
		if detaillen.Len > logmodel["location"].Length {
			return errors.New("location too long," + fmt.Sprint(detaillen.Len) + "/" + fmt.Sprint(logmodel["location"].Length))
		}
	}

	tx := cli.Table("log").Create(log)
	err = tx.Error
	if err == nil {
		xlog.Logger.Info("insert log", zap.Any("log", log))
	}
	return
}

func GetLogs(start, end time.Time, quadrant int, location, atype, title, detail, review string, offset, limit int, orderby string) (bs []model.Log, total int64, err error) {
	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	tx := cli.Table("log")
	txtotal := cli.Table("log")

	if quadrant > 0 {
		tx = tx.Where("quadrant & ? = ?", quadrant, quadrant)
		txtotal = txtotal.Where("quadrant & ? = ?", quadrant, quadrant)
	}
	if start.Unix() != (time.Time{}.Unix()) {
		tx = tx.Where("ctime >= ?", start)
		txtotal = txtotal.Where("ctime >= ?", start)
	}
	if end.Unix() != (time.Time{}.Unix()) {
		tx = tx.Where("ctime <= ?", end)
		txtotal = txtotal.Where("ctime <= ?", end)
	}
	if location != "" {
		//"name LIKE ?", "%jin%"
		tx = tx.Where("location like ?", "%"+location+"%")
		txtotal = txtotal.Where("location like ?", "%"+location+"%")
	}
	if atype != "" {
		tx = tx.Where("atype like ?", "%"+atype+"%")
		txtotal = txtotal.Where("atype like ?", "%"+atype+"%")
	}
	if title != "" {
		tx = tx.Where("title like ?", "%"+title+"%")
		txtotal = txtotal.Where("title like ?", "%"+title+"%")
	}
	if detail != "" {
		tx = tx.Where("detail like ?", "%"+detail+"%")
		txtotal = txtotal.Where("detail like ?", "%"+detail+"%")
	}
	if review != "" {
		tx = tx.Where("review like ?", "%"+review+"%")
		txtotal = txtotal.Where("review like ?", "%"+review+"%")
	}

	tx = tx.Offset(offset).Limit(limit)
	if orderby != "" {
		tx = tx.Order(orderby)
	} else {
		tx = tx.Order("ctime desc")
	}
	err = tx.Find(&bs).Error

	if err != nil {
		panic(err)
		xlog.Logger.Error("sql", zap.Error(err))
		return
	}

	err = txtotal.Count(&total).Error
	if err != nil {
		panic(err)
		xlog.Logger.Error("sql", zap.Error(err))
		return
	}
	return
}
