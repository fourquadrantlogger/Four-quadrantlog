package service

import (
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func CreateBlob(log *model.Blob) (err error) {
	log.ID = uuid.New().String()

	_, cli, err := mysqlcli.GetMysqlClient()
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

func GetBlob(id string) (b model.Blob, err error) {

	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	tx := cli.Table("log").Where("id=?", id).First(&b)
	err = tx.Error

	return
}
