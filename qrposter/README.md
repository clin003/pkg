#go-pkg qrposter

##	介绍

二维码合成海报工具包

##	使用
go get  gitee.com/lyhuilin/pkg

###	例子

在example文件夹中有简单的使用例子

···golang

	package main
	
	import (
		"os"
		"path"
	
		"gitee.com/lyhuilin/demo/gogenposter/pkg/qrposter"
	
		"github.com/boombuler/barcode/qr"
		// "github.com/spf13/viper"
	)
	
	func main() {
		Generate()
	}
	func Generate() (err error) {
		url := "https://www.lyhuilin.com"
		dirDst := "./data/gen/dst"
		dirQrcode := "./data/gen/qrcode"
	
		os.MkdirAll(dirDst, os.ModePerm)
		os.MkdirAll(dirQrcode, os.ModePerm)
	
		qrWidth := 250
		qrHeight := 250
		// 生成二维码
		qrcode := qrposter.NewQrCode(url, qrWidth, qrHeight, qr.M, qr.Auto)
		filePath, err := qrcode.Encode(dirQrcode)
		if err != nil {
			panic(err)
		}
	
		bgPath := "./data/img/poster1.jpg"
		dstPath := path.Join(dirDst, qrcode.FileName)
	
		rectX0 := 0
		rectY0 := 0
		rectX1 := 750
		rectY1 := 1334
	
		qrX := 250
		qrY := 978
	
		poster := qrposter.NewPoster(
			qrposter.Content{
				BgPath:  bgPath,
				DstPath: dstPath,
			},
			&qrposter.Rect{
				X0: rectX0,
				Y0: rectY0,
				X1: rectX1,
				Y1: rectY1,
			},
			qrposter.Qr{
				Path: filePath,
				X:    qrX,
				Y:    qrY,
			},
		)
	
		err = poster.Generate()
		return
	}
···