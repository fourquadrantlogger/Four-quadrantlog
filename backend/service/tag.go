package service

import (
	"fmt"
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/storage/mysqlcli"
	"go.uber.org/zap"
	"time"
)

type Tag struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func GroupbyTags(start, end time.Time, quadrant int, location, atype, title, detail, review string, limit int) (bs []Tag, err error) {
	_, tx, err := mysqlcli.GetMysqlClient()
	if err != nil {
		return
	}
	s := `
select tag as name, count( tag)as value from log,
json_table(
  atype,
  '$[*]' columns(
     tag varchar(128) path '$'
    ) 
) as tag
`

	swhere := "where "
	if quadrant > 0 {
		swhere = swhere + (" quadrant = " + fmt.Sprint(quadrant) + " and ")
	}
	if start.Unix() != (time.Time{}.Unix()) {
		swhere = swhere + (" ctime  >= '" + fmt.Sprint(start) + "' and ")
	}
	if end.Unix() != (time.Time{}.Unix()) {

		swhere = swhere + (" ctime <= '" + fmt.Sprint(end) + "' and ")
	}
	if location != "" {
		swhere = swhere + (" location like  '%" + location + "%' and ")
	}
	if atype != "" {
		swhere = swhere + ("atype like  '%" + atype + "%'" + " and ")
	}
	if title != "" {
		swhere = swhere + (" title like  '%" + title + "%'" + " and ")
	}
	if detail != "" {
		swhere = swhere + (" detail like '%" + detail + "%' and ")
	}
	if review != "" {
		swhere = swhere + (" review like '%" + location + "%' and ")
	}
	swhere += " 1=1"

	s += swhere + " group by tag  order by value desc limit " + fmt.Sprint(limit)

	err = tx.Raw(s).Scan(&bs).Error

	if err != nil {
		xlog.Logger.Error("sql", zap.Error(err))
		return
	}

	return
}
