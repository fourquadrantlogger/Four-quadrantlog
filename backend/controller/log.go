package controller

import (
	"errors"
	"fourquadrantlog/assist/xtime"
	"fourquadrantlog/model"
	"fourquadrantlog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type CommonQuery struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func GetLogs(c *gin.Context) {
	var start, end time.Time
	var offset, limit = 0, 20
	if c.Query("start") != "" {
		startStr, _ := url.QueryUnescape(c.Query("start"))
		o, err := xtime.Parse(startStr)
		if err == nil {
			start = o
		}
	}
	if c.Query("end") != "" {
		endStr, _ := url.QueryUnescape(c.Query("end"))
		o, err := xtime.Parse(endStr)
		if err == nil {
			end = o
		}
	}
	if c.Query("offset") != "" {
		o, _ := strconv.Atoi(c.Query("offset"))
		if o != 0 {
			offset = o
		}
	}
	if c.Query("limit") != "" {
		o, _ := strconv.Atoi(c.Query("limit"))
		if o != 0 {
			limit = o
		}
	}
	var quadrantint int = model.Quadrant2Int(c.Query("quadrant"))
	b, total, err := service.GetLogs(start, end,
		quadrantint,
		c.Query("location"),
		c.Query("atype"),
		c.Query("title"),
		c.Query("detail"),
		c.Query("review"),
		offset,
		limit,
		c.Query("orderby"),
	)
	if err != nil {
		c.Error(err)
	} else {
		for i, _ := range b {
			b[i].FixShow()
			if b[i].Detail != nil && len(*b[i].Detail) > 500 {
				var detail = (*b[i].Detail)[:500]
				b[i].Detail = &detail
			}

			if b[i].Review != nil && len(*b[i].Review) > 500 {
				var review = (*b[i].Review)[:500]
				b[i].Review = &review
			}
		}
		c.JSON(200, CommonQuery{
			Data:  b,
			Total: total,
		})
	}
}
func CreateLog(c *gin.Context) {
	var log model.Log
	if c.Bind(&log) == nil {
		if log.Ctime_ != "" {
			var e error
			log.Ctime, e = xtime.Parse(log.Ctime_)
			if e != nil {
				log.Ctime = time.Now()
			}
		}

		log.FixForDB()
		err := service.CreateLog(&log)
		if err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   log,
			})
		}
	}
}

func UpdateLog(c *gin.Context) {
	var log model.Log

	if c.Bind(&log) == nil {
		if log.ID == "" {
			c.Error(errors.New("need id"))
			return
		}
		if log.Ctime_ != "" {
			log.Ctime, _ = time.Parse("2006-01-02 15:04:05", log.Ctime_)
		}
		if log.Ctime.Unix() == (time.Time{}.Unix()) {
			log.Ctime = time.Now()
		}
		log.FixForDB()
		err := service.UpdateLog(&log)
		if err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   log,
			})
		}
	}
}

func DeleteLog(c *gin.Context) {
	logid := c.Params.ByName("logid")
	if logid == "" {
		c.Error(errors.New("id not set"))
		return
	}
	err := service.DeleteLog(logid)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   logid,
		})
	}

}
