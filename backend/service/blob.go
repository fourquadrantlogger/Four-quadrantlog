package service

import (
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
)

func GetBlob(id string, selectblob bool) (b model.Log, err error) {
	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	selects := "id,title,quadrant,atype,detail,review,location,ctime"
	if selectblob {
		selects += ",blob"
	}
	tx := cli.Table("log").Select(selects).Where("id=?", id).First(&b)
	err = tx.Error

	return
}
