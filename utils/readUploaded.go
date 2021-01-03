package utils

import (
	"io/ioutil"
	"mime/multipart"
)

func ReadUploaded(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	return ioutil.ReadAll(src)
}
