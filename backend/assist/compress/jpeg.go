package compress

import (
	"bytes"
	"fourquadrantlog/assist/xlog"
	"go.uber.org/zap"
	"image/jpeg"
)

func CompressJpeg(imgraw []byte,quality int)(compressed []byte,err error){
	reader:=bytes.NewReader(imgraw)
	jpgimg, err := jpeg.Decode(reader)
	if err != nil {
		xlog.Logger.Warn("jpeg.Decode::", zap.Error(err))
		return
	}

	//保存到新文件中
	newfile :=  bytes.NewBuffer(imgraw)
	if err != nil {
		return
	}


	// &jpeg.Options{Quality: 10} 图片压缩质量
	err = jpeg.Encode(newfile, jpgimg, &jpeg.Options{Quality: quality})
	if err != nil {
		xlog.Logger.Warn("jpeg.Encode::", zap.Error(err))
		return
	}
	compressed=newfile.Bytes()
	return
}
