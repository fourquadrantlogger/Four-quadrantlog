package service

import (
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
)

func GetBlob(id string) (b model.Log, err error) {

	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	tx := cli.Table("log").Where("id=?", id).First(&b)
	err = tx.Error

	return
}
