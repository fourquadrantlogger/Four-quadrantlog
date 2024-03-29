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
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetLog(c *gin.Context) {
	logid := c.Params.ByName("logid")

	b, err := service.GetBlob(logid, false)
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

	b, err := service.GetBlob(blobid, true)

	if err != nil {
		c.Error(err)
	} else {
		var contentType string
		if filepath.Ext(b.Title) != "" {
			contentType = mime.TypeByExtension(filepath.Ext(b.Title))
		} else {
			contentType = http.DetectContentType(b.Blob)
		}

		b.FixShow()
		c.Data(200, contentType, b.Blob)
	}
}

func GetCompressed(c *gin.Context) {
	blobid := c.Params.ByName("blobid")

	b, err := service.GetBlob(blobid, true)
	if err != nil {
		c.Error(err)
		return
	}
	ext := strings.ToLower(path.Ext(b.Title))
	quality := 30
	if c.Query("quality") != "" {
		quality, _ = strconv.Atoi(c.Query("quality"))
	}
	if ext == ".jpg" || ext == ".jpeg" {
		b.Blob, err = compress.CompressJpeg(b.Blob, quality)
		if err != nil {
			c.Error(err)
		} else {
			contentType := "image/jpeg"
			c.Data(200, contentType, b.Blob)
		}
	}
	if ext == ".png" {
		b.Blob, err = compress.Png2Jpeg(b.Blob, quality)
		if err != nil {
			c.Error(err)
		} else {
			contentType := "image/jpeg"
			c.Data(200, contentType, b.Blob)
		}
	}
}

func CreateBlob(c *gin.Context) {
	var log model.Log

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
	log.Blob, err = io.ReadAll(f)
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	err = service.CreateLog(&log)
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
