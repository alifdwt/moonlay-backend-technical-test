package util

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Read file extension, to make sure it's an pdf of txt
	fileExt := filepath.Ext(file.Filename)
	if fileExt != ".pdf" && fileExt != ".txt" {
		return "", errors.New("file must be pdf or txt")
	}

	// Destination path to upload the file
	dstPath := "uploads/"
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		err = os.MkdirAll(dstPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	fileName := filepath.Join(dstPath, file.Filename)
	dst, err := os.Create(fileName)

	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = dst.ReadFrom(src); err != nil {
		return "", err
	}

	return fileName, nil
}