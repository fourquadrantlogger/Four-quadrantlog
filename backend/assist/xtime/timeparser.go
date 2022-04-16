package xtime

import (
	"errors"
	"time"
)

var Tformat = map[string]string{
	"0.":   "2006.01.02 15:04:05",
	".":    "2006.1.2 15:04:05",
	"0-":   "2006-01-02 15:04:05",
	"-":    "2006-1-2 15:04:05",
	"0/":   "2006/1/2 15:04:05",
	"/":    "2006/01/02 15:04:05",
	"/day": "2006/01/02",
	"-day": "2006/1/2",
}

func Parse(t string) (r time.Time, e error) {
	for _, f := range Tformat {
		var err error
		r, err = time.Parse(f, t)
		if err == nil {
			return
		}
	}
	return r, errors.New("format not found:" + t)
}
