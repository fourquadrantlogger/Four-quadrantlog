package model

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	ID        string    `json:"id"`
	Quadrant  int       `json:"-"`
	Quadrant_ string    `json:"quadrant" gorm:"-"`
	Ctime     time.Time `json:"-"`
	Ctime_    string    `json:"ctime" gorm:"-"`
	Location  *string   `json:"location"`
	Atype     string    `json:"atype" gorm:"atype"`
	Title     string    `json:"title"`
	Detail    *string   `json:"detail"`
	Review    *string   `json:"review"`
}

func (l *Log) FixShow() {
	l.Quadrant_ = strings.Join(Int2Quadrant(l.Quadrant), "/")
	var atype_ []string
	json.Unmarshal([]byte(l.Atype), &atype_)
	l.Atype = strings.Join(atype_, "/")
	l.Ctime_ = l.Ctime.Format("2006-01-02 15:04:05")
}

func (l *Log) FixForDB() {
	l.Quadrant = Quadrant2Int(l.Quadrant_)

	var atype_ = strings.Split(l.Atype, "/")
	AtypeBytes, _ := json.Marshal(atype_)
	l.Atype = string(AtypeBytes)

	l.Ctime_ = l.Ctime.Format("2006-01-02 15:04:05")
}
func Quadrant2Int(S string) (V int) {
	var a [5]bool
	ss := strings.Split(S, "/")
	for _, s := range ss {
		if s == "时空" {
			a[0] = true
		}
		if s == "肉体" {
			a[1] = true
		}
		if strings.Contains(s, "信息") || strings.Contains(s, "知识") {
			a[2] = true
		}
		if s == "社会关系" {
			a[3] = true
		}
		if s == "持有物" {
			a[4] = true
		}
	}

	for i, b := range a {
		var v = 1
		if b {
			v = v << i
			V = V | v
		}
	}
	return
}
func Int2Quadrant(V int) (ss []string) {

	vs := strconv.FormatInt(int64(V), 2)
	for i := len(vs); i < 5; i++ {
		vs = "0" + vs
	}
	for i, s := range vs {
		if s == '1' {
			switch 4 - i {
			case 0:
				ss = append(ss, "时空")
			case 1:
				ss = append(ss, "肉体")
			case 2:
				ss = append(ss, "信息知识")
			case 3:
				ss = append(ss, "社会关系")
			case 4:
				ss = append(ss, "持有物")
			}
		}
	}
	return
}
