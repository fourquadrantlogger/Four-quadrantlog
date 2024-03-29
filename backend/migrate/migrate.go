package main

import (
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/model"
	"fourquadrantlog/storage/mysqlcli"
	"go.uber.org/zap"
	"os"
)

func main() {

	local, err := mysqlcli.MysqlClient("root:Abc_014916@tcp(rm-j6c8z8i3llo2r369vbo.mysql.rds.aliyuncs.com:3306)/fourquadrantlog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		xlog.Logger.Debug("", zap.Error(err))
		return
	}

	remote, err := mysqlcli.MysqlClient(os.Getenv("MYSQL"))
	if err != nil {
		xlog.Logger.Debug("", zap.Error(err))
		return
	}

	rows, err := local.Table("log").Order("ctime").Select("*").Rows() // (*sql.Rows, error)
	defer rows.Close()

	var blog model.Blob
	for rows.Next() {
		// ScanRows scan a row into user
		err = local.ScanRows(rows, &blog)
		if err != nil {
			xlog.Logger.Error("x", zap.Error(err))
			panic(err)
			break
		}

		err = remote.Table("log").Create(blog).Error
		if err != nil {
			xlog.Logger.Error("x", zap.Error(err))
			break
		}
	}

}
