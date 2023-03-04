package service

import (
	"errors"
	"fmt"
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
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

	cli.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if e := tx.Table("log").Create(log).Error; e != nil {
			// 返回任何错误都会回滚事务
			return e
		}
		for _, t := range log.Atype_ {
			tag := model.Tag{
				Log: log.ID,
				Tag: t,
			}
			if e := tx.Table("tag").Create(tag).Error; e != nil {
				return e
			}
			xlog.Logger.Info("insert tag", zap.Any("tag", tag))
		}

		// 返回 nil 提交事务
		return nil
	})

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

	var quotacheck = func(field string) string {
		runefiled := []rune(field)
		if len(runefiled) > 2 && runefiled[0] == '\'' && field[len(runefiled)-1] == '\'' {
			return string(runefiled[1 : len(runefiled)-1])
		}
		return field
	}
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
		atypes := strings.Split(atype, ",")
		tx.Joins("JOIN tag ON tag.log = log.id ")
		txtotal.Joins("left JOIN tag ON tag.log = log.id ")
		for _, at := range atypes {

			shortat := quotacheck(at)

			if shortat != at {
				tx = tx.Where("tag.tag= ?", shortat)
				txtotal = txtotal.Where("tag.tag= ?", shortat)
			} else {
				tx = tx.Where("tag.tag like ?", "%"+shortat+"%")
				txtotal = txtotal.Where("tag.tag like ?", "%"+shortat+"%")
			}

		}

	}
	if title != "" {
		at := quotacheck(title)
		if at != title {
			tx = tx.Where("title = ?", at)
			txtotal = txtotal.Where("title = ?", at)
		} else {
			tx = tx.Where("title like ?", "%"+title+"%")
			txtotal = txtotal.Where("title like ?", "%"+title+"%")
		}

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
