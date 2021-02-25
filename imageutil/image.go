package imageutil

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
)

type ImageInfo struct {
	Content []byte
	ExtName string
	Height  int
	Width   int
}

func ImageResizeChange(source []byte, quality, width int) (*ImageInfo, error) {
	if len(source) == 0 {
		return nil, errors.New("ImageResizeChange Picture Source is nil.")
	}

	config, suffix, err := image.DecodeConfig(bytes.NewBuffer(source))
	if err != nil {
		return nil, err
	}

	var origin image.Image

	switch strings.ToLower(suffix) {
	case "jpg", "jpeg":

		origin, err = jpeg.Decode(bytes.NewReader(source))
		if err != nil {
			return nil, err
		}

	case "png":

		origin, err = png.Decode(bytes.NewReader(source))
		if err != nil {
			return nil, err
		}

	case "GIF", "gif":
		origin, err = gif.Decode(bytes.NewReader(source))
		if err != nil {
			return nil, err
		}

	case "webp":
		origin, err = webp.Decode(bytes.NewReader(source))
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("ImageResizeChange Read Image type error:" + suffix)
	}

	if width <= 0 {
		width = config.Width
	}

	if width >= config.Width {
		width = config.Width
	}
	widthTemp := uint(width)
	//等比缩放
	height := uint(width * config.Height / config.Width)
	canvas := resize.Thumbnail(widthTemp, height, origin, resize.Lanczos3)

	tempFile, fileErr := ioutil.TempFile("", "primas_size_image")
	if fileErr != nil {
		err = fileErr
		return nil, fileErr
	}
	defer func() {
		_ = os.Remove(tempFile.Name())
		_ = tempFile.Close()
	}()

	switch strings.ToLower(suffix) {
	case "jpg", "jpeg":
		err = jpeg.Encode(tempFile, canvas, &jpeg.Options{Quality: quality})
		if err != nil {
			return nil, err
		}

	case "png":
		err = png.Encode(tempFile, canvas)
		if err != nil {
			return nil, err
		}

	case "GIF", "gif":
		err = gif.Encode(tempFile, canvas, nil)
		if err != nil {
			return nil, err
		}

	case "webp":
		err = png.Encode(tempFile, canvas)
		if err != nil {
			return nil, err
		}
		suffix = "png"

	default:
		return nil, errors.New("ImageResizeChange Write Image type error.")
	}

	_, _ = tempFile.Seek(0, 0)
	target, err := ioutil.ReadAll(tempFile)
	if err != nil {
		return nil, err
	}

	targetImage := ImageInfo{
		Content: target,
		ExtName: suffix,
		Height:  config.Height,
		Width:   config.Width,
	}

	return &targetImage, nil
}

func ImageBaseInfo(source []byte, quality, width int) (*ImageInfo, error) {
	if len(source) == 0 {
		return nil, errors.New("ImageResizeChange Picture Source is nil.")
	}

	config, suffix, err := image.DecodeConfig(bytes.NewBuffer(source))
	if err != nil {
		return nil, err
	}

	target := source

	switch strings.ToLower(suffix) {
	case "jpg", "jpeg":

	case "png":

	case "GIF", "gif":

	case "webp":
		var origin image.Image

		if width <= 0 {
			width = config.Width
		}

		if width >= config.Width {
			width = config.Width
		}
		widthTemp := uint(width)
		//等比缩放
		height := uint(width * config.Height / config.Width)
		canvas := resize.Thumbnail(widthTemp, height, origin, resize.Lanczos3)

		tempFile, fileErr := ioutil.TempFile("", "primas_size_image")
		if fileErr != nil {
			err = fileErr
			return nil, fileErr
		}
		defer func() {
			_ = os.Remove(tempFile.Name())
			_ = tempFile.Close()
		}()

		err = png.Encode(tempFile, canvas)
		if err != nil {
			return nil, err
		}
		suffix = "png"

		origin, err = webp.Decode(bytes.NewReader(source))
		if err != nil {
			return nil, err
		}

		_, _ = tempFile.Seek(0, 0)
		target, err = ioutil.ReadAll(tempFile)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("ImageResizeChange Read Image type error:" + suffix)
	}

	targetImage := ImageInfo{
		Content: target,
		ExtName: suffix,
		Height:  config.Height,
		Width:   config.Width,
	}

	return &targetImage, nil
}
