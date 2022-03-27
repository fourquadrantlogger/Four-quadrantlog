package controller

import (
	"fmt"
	"fourquadrantlog/model"
	"fourquadrantlog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetLogs(c *gin.Context) {
	var start, end time.Time
	var offset, limit = 0, 20
	if c.Query("start") != "" {
		o, _ := time.Parse("2006-01-02 15:04:05", c.Query("start"))
		if o.Unix() != 0 {
			start = o
		} else {
			o, _ = time.Parse("2006-01-02", c.Query("start"))
			if o.Unix() != 0 {
				start = o
			}
		}

	}
	if c.Query("end") != "" {
		o, _ := time.Parse("2006-01-02 15:04:05", c.Query("end"))
		if o.Unix() != 0 {
			end = o
		} else {
			o, _ = time.Parse("2006-01-02", c.Query("end"))
			if o.Unix() != 0 {
				end = o
			}
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
	b, err := service.GetLogs(start, end,
		c.Query("location"),
		c.Query("quadrant"),
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
		c.JSON(200, b)
	}
}
func CreateLog(c *gin.Context) {
	var log model.Log
	if c.Bind(&log) == nil {
		fmt.Println(log.Ctime.Unix())
		if log.Ctime.Unix() == (time.Time{}.Unix()) {
			log.Ctime = time.Now()
		}

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
