package main

import (
	"os"
	"path"

	"gitee.com/lyhuilin/pkg/qrposter"

	"github.com/boombuler/barcode/qr"
	// "github.com/spf13/viper"
)

func main() {
	Generate()
}
func Generate() (err error) {
	url := "https://www.lyhuilin.com"
	// 生成内容存储目录
	dirDst := "./data/gen/dst"
	// 二维码存储目录
	dirQrcode := "./data/gen/qrcode"

	os.MkdirAll(dirDst, os.ModePerm)
	os.MkdirAll(dirQrcode, os.ModePerm)

	// 二维码宽高
	qrWidth := 250
	qrHeight := 250
	// 生成二维码
	qrcode := qrposter.NewQrCode(url, qrWidth, qrHeight, qr.M, qr.Auto)
	filePath, err := qrcode.Encode(dirQrcode)
	if err != nil {
		panic(err)
	}

	// 海报背景图片地址
	bgPath := "./data/img/poster1.jpg"
	dstPath := path.Join(dirDst, qrcode.FileName)

	// 画布图形设置
	rectX0 := 0
	rectY0 := 0
	rectX1 := 750
	rectY1 := 1334

	// 二维码位置
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

// func Generate() (err error) {
// 	url := viper.GetString("qr_poster_generate.url")
// 	dirDst := viper.GetString("qr_poster_generate.dst_dir")
// 	dirQrcode := viper.GetString("qr_poster_generate.qrcode_dir")

// 	os.MkdirAll(dirDst, os.ModePerm)
// 	os.MkdirAll(dirQrcode, os.ModePerm)

// 	qrWidth := viper.GetInt("qr_poster_generate.qrcode.width")
// 	qrHeight := viper.GetInt("qr_poster_generate.qrcode.height")
// 	// 生成二维码
// 	qrcode := qrposter.NewQrCode(url, qrWidth, qrHeight, qr.M, qr.Auto)
// 	filePath, err := qrcode.Encode(dirQrcode)
// 	if err != nil {
// 		panic(err)
// 	}

// 	bgPath := viper.GetString("qr_poster_generate.content.bg_path")
// 	dstPath := path.Join(dirDst, qrcode.FileName)

// 	rectX0 := viper.GetInt("qr_poster_generate.rect.x0")
// 	rectY0 := viper.GetInt("qr_poster_generate.rect.y0")
// 	rectX1 := viper.GetInt("qr_poster_generate.rect.x1")
// 	rectY1 := viper.GetInt("qr_poster_generate.rect.y1")

// 	qrX := viper.GetInt("qr_poster_generate.qrcode.x0")
// 	qrY := viper.GetInt("qr_poster_generate.qrcode.y0")

// 	poster := qrposter.NewPoster(
// 		qrposter.Content{
// 			BgPath:  bgPath,
// 			DstPath: dstPath,
// 		},
// 		&qrposter.Rect{
// 			X0: rectX0,
// 			Y0: rectY0,
// 			X1: rectX1,
// 			Y1: rectY1,
// 		},
// 		qrposter.Qr{
// 			Path: filePath,
// 			X:    qrX,
// 			Y:    qrY,
// 		},
// 	)

// 	err = poster.Generate()
// 	return
// }
