package controller

import (
	"fmt"
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/model"
	"fourquadrantlog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"time"
)

func GetLog(c *gin.Context) {
	_logid := c.Params.ByName("logid")
	logid, _ := strconv.Atoi(_logid)
	b, err := service.GetBlob(logid)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(200, b)
	}
}

func GetBlob(c *gin.Context) {
	_blobid := c.Params.ByName("blobid")
	blobid, _ := strconv.Atoi(_blobid)
	b, err := service.GetBlob(blobid)
	if err != nil {
		c.Error(err)
	} else {
		contentType := http.DetectContentType(b.Blob)
		c.Data(200, contentType, b.Blob)
	}
}

func CreateBlob(c *gin.Context) {
	var log model.Blob

	log.Atype = c.PostForm("atype")
	log.Quadrant = c.PostForm("quadrant")
	var err error
	log.Ctime, err = time.Parse("2006-01-02 15:04:05", c.PostForm("ctime"))
	if err != nil {
		xlog.Logger.Warn("ctime err,will use time.now", zap.Error(err))
	} else {
		log.Ctime = time.Now()
	}

	lo := c.PostForm("location")
	if lo != "" {
		log.Location = &lo
	}
	re := c.PostForm("review")
	if re != "" {
		log.Review = &re
	}

	de := c.PostForm("detail")
	if de != "" {
		log.Detail = &de
	}

	// Source

	file, err := c.FormFile("blob")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	if c.PostForm("blobrename") != "" {

		log.Title = c.PostForm("blobrename") + path.Ext(file.Filename)
	} else {
		log.Title = file.Filename
	}

	f, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	log.Blob, err = ioutil.ReadAll(f)
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	err = service.CreateBlob(&log)
	if err != nil {
		c.Error(err)
		return
	} else {
		log.Blob = []byte(fmt.Sprint(len(log.Blob)))
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   log,
		})
	}
}
