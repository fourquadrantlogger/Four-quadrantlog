package controller

import (
	"encoding/json"
	"fmt"
	"fourquadrantlog/assist/compress"
	"fourquadrantlog/assist/xlog"
	"fourquadrantlog/assist/xtime"
	"fourquadrantlog/model"
	"fourquadrantlog/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
)

func GetLog(c *gin.Context) {
	logid := c.Params.ByName("logid")

	b, err := service.GetBlob(logid)
	b.Blob = nil
	if err != nil {
		c.Error(err)
	} else {
		b.FixShow()
		c.JSON(200, b)
	}
}

func GetBlob(c *gin.Context) {
	blobid := c.Params.ByName("blobid")

	b, err := service.GetBlob(blobid)

	if err != nil {
		c.Error(err)
	} else {
		contentType := http.DetectContentType(b.Blob)
		b.FixShow()
		c.Data(200, contentType, b.Blob)
	}
}

func GetCompressed(c *gin.Context) {
	blobid := c.Params.ByName("blobid")

	b, err := service.GetBlob(blobid)
	if err != nil {
		c.Error(err)
		return
	}
	ext:=strings.ToLower(path.Ext(b.Title))
	if ext==".jpg"||ext==".jpeg"{
		b.Blob,err=compress.CompressJpeg(b.Blob,30)
		if err != nil {
			c.Error(err)
		} else {
			contentType := http.DetectContentType(b.Blob)
			b.FixShow()
			c.Data(200, contentType, b.Blob)
		}
	}

}

func CreateBlob(c *gin.Context) {
	var log model.Blob

	var Atype_ = strings.Split(c.PostForm("atype"), "/")
	atypes, _ := json.Marshal(Atype_)
	log.Atype = string(atypes)

	log.Quadrant_ = c.PostForm("quadrant")
	log.Quadrant = model.Quadrant2Int(log.Quadrant_)
	var err error
	log.Ctime, err = xtime.Parse(c.PostForm("ctime"))
	if err != nil {
		xlog.Logger.Warn("ctime err,will use time.now", zap.Error(err))
		log.Ctime = time.Now()
		err = nil
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
