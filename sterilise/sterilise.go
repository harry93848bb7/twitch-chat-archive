package sterilise

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// UnknownFormat ...
var UnknownFormat = errors.New("Unknown image format")

// SteriliseImage ...
func SteriliseImage(b []byte) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		return nil, "", UnknownFormat
	}
	switch format {
	case "png":
		var buffer bytes.Buffer
		if err := png.Encode(&buffer, img); err != nil {
			return nil, "", err
		}
		return buffer.Bytes(), format, nil
	case "jpeg":
		var buffer bytes.Buffer
		if err := jpeg.Encode(&buffer, img, &jpeg.Options{Quality: 100}); err != nil {
			return nil, "", err
		}
		return buffer.Bytes(), format, nil
	case "gif":
		var buffer bytes.Buffer
		g, err := gif.DecodeAll(bytes.NewBuffer(b))
		if err != nil {
			return nil, "", err
		}
		if err := gif.EncodeAll(&buffer, g); err != nil {
			return nil, "", err
		}
		return buffer.Bytes(), format, nil
	default:
		return nil, "", UnknownFormat
	}
}
