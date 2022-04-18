package service

import (
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/storage/mysqlcli"
	"go.uber.org/zap"
)

var models = make(map[string]map[string]Dbmodel)

func Model(db, model string) map[string]Dbmodel {

	if m, h := models[model]; h {
		return m
	}
	m, err := GetModel(db, model)
	if err != nil {
		xlog.Logger.Error("get model err", zap.Error(err))
	} else {
		models[model] = m
	}
	return m
}

type Dbmodel struct {
	TABLE_NAME     string `gorm:"TABLE_NAME"`
	COLUMN_COMMENT string `gorm:"COLUMN_COMMENT"`
	COLUMN_NAME    string `gorm:"COLUMN_NAME"`
	DATA_TYPE      string `gorm:"DATA_TYPE"`
	Length         int    `gorm:"length"`
	AutoIncrement  bool   `gorm:"auto_increment"`
	IsNullable     bool   `gorm:"IS_NULLABLE"`
	COLUMN_KEY     bool   `gorm:"COLUMN_KEY"`
	COLUMN_DEFAULT string `gorm:"COLUMN_DEFAULT"`
}

func GetModel(db, table string) (model map[string]Dbmodel, err error) {
	var dbfields = make([]Dbmodel, 0)
	model = make(map[string]Dbmodel)
	sql := `
	
SELECT 
  
    COLUMN_COMMENT,
    Column_Name,
    data_type,
    (CASE
        WHEN
            data_type = 'float'
                OR data_type = 'double'
                OR data_type = 'TINYINT'
                OR data_type = 'SMALLINT'
                OR data_type = 'MEDIUMINT'
                OR data_type = 'INT'
                OR data_type = 'INTEGER'
                OR data_type = 'decimal'
                OR data_type = 'bigint'
        THEN
            NUMERIC_PRECISION
        ELSE CHARACTER_MAXIMUM_LENGTH
    END) AS 'length',
    NUMERIC_SCALE,
    (CASE
        WHEN EXTRA = 'auto_increment' THEN 1
        ELSE 0
    END) AS 'auto_increment',
    (CASE
        WHEN IS_NULLABLE = 'NO' THEN 0
        ELSE 1
    END) AS 'IS_NULLABLE',
    (CASE
        WHEN COLUMN_KEY = 'PRI' THEN 1
        ELSE 0
    END) AS 'COLUMN_KEY',
    COLUMN_DEFAULT AS 'COLUMN_DEFAULT'
FROM
    information_schema.COLUMNS
WHERE
    table_schema = '` + db + `' and   table_name = '` + table + `' 

`
	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	err = cli.Raw(sql).Scan(&dbfields).Error

	for _, f := range dbfields {
		model[f.COLUMN_NAME] = f
	}
	return
}

type LenModel struct {
	Len int
}

func CheckStringLen(value string) (l LenModel, err error) {
	sql := `SELECT 
       length( ? )  as len`
	_, cli, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	err = cli.Raw(sql, value).Scan(&l).Error
	return
}
