package qrcode

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(url string, size int) (image.Image, error) {
	code, err := qrcode.New(url, qrcode.Highest)
	if err != nil {
		return nil, err
	}
	return code.Image(size), nil
}

func createJPEGFile(img image.Image, imgName string) error {
	outFile, err := os.Create(imgName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	return nil
}

func SaveQRCode(url string, size int, fileName string) error {
	qr, err := generateQRCode(url, size)
	if err != nil {
		return nil
	}
	return createJPEGFile(qr, fileName)
}
