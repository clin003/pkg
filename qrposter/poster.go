package qrposter

import (
	"image"
	"image/draw"
	"image/jpeg"

	// "io/ioutil"
	"os"
	// "github.com/golang/freetype"
)

type Poster struct {
	*Content
	*Rect
	Qr *Qr
}

type Content struct {
	BgPath  string
	DstPath string
	DstFile *os.File
}

type Rect struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

type Qr struct {
	Path string
	X    int
	Y    int
}

func NewPoster(content Content, rect *Rect, qr Qr) *Poster {
	return &Poster{
		Content: &content,
		Rect:    rect,
		Qr:      &qr,
	}
}

func (p *Poster) Generate() (err error) {
	p.DstFile, err = os.OpenFile(p.DstPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	bgFile, err := os.Open(p.BgPath)
	if err != nil {
		return err
	}
	defer bgFile.Close()

	bgImage, err := jpeg.Decode(bgFile)
	if err != nil {
		return err
	}

	qrFile, err := os.Open(p.Qr.Path)
	if err != nil {
		return err
	}
	defer qrFile.Close()

	qrImage, err := jpeg.Decode(qrFile)
	if err != nil {
		return err
	}

	jpg := image.NewRGBA(image.Rect(p.Rect.X0, p.Rect.Y0, p.Rect.X1, p.Rect.Y1))
	draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
	draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(p.Qr.X, p.Qr.Y)), draw.Over)

	err = jpeg.Encode(p.DstFile, jpg, nil)
	if err != nil {
		return err
	}

	return nil
}
