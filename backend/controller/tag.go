package controller

import (
	"fourquadrantlog/assist/xtime"
	"fourquadrantlog/model"
	"fourquadrantlog/service"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
	"time"
)

func GetTags(c *gin.Context) {
	var start, end time.Time
	var limit = 20
	if c.Query("limit") != "" {
		limit, _ = strconv.Atoi(c.Query("limit"))
	}
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
	var quadrantint int = model.Quadrant2Int(c.Query("quadrant"))
	b, err := service.GroupbyTags(start, end,
		quadrantint,
		c.Query("location"),
		c.Query("title"),
		c.Query("detail"),
		c.Query("review"),
		limit,
	)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(200, b)
	}
}
