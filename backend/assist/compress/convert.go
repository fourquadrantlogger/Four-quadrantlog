package compress

import (
	"bytes"
	"fourquadrantlog/assist/xlog"
	"go.uber.org/zap"
	"image/jpeg"
	"image/png"
)

func Png2Jpeg(imgraw []byte, quality int) (compressed []byte, err error) {
	reader := bytes.NewReader(imgraw)
	img, err := png.Decode(reader)
	if err != nil {
		xlog.Logger.Warn("png.Decode::", zap.Error(err))
		return
	}

	//保存到新文件中

	// 转jpg

	newfile := bytes.NewBuffer(compressed)
	if err != nil {
		return
	}
	err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: quality})
	if err != nil {
		xlog.Logger.Warn("jpeg.Encode::", zap.Error(err))
		return
	}
	compressed = newfile.Bytes()
	return
}
