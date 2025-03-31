package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func TrainCaptcha(uid int32, data []byte, captcha string) (string, string, error) {
	dir := filepath.Join("upload", fmt.Sprintf("%d", uid))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", "", err
	}

	// Tạo tên file dựa trên timestamp
	filename := time.Now().UnixNano()
	filenameImg := fmt.Sprintf("%d.png", filename)
	filenameCaptcha := fmt.Sprintf("%d.txt", filename)
	filePathImg := filepath.Join(dir, filenameImg)
	filePathCaptcha := filepath.Join(dir, filenameCaptcha)

	// Ghi dữ liệu vào file
	err = os.WriteFile(filePathImg, data, 0644)
	if err != nil {
		return "", "", err
	}

	err = os.WriteFile(filePathCaptcha, []byte(captcha), 0644)
	if err != nil {
		return "", "", err
	}

	return filePathImg, filePathCaptcha, nil
}

func SaveImage(uid int32, data []byte) (string, error) {
	dir := filepath.Join("upload", fmt.Sprintf("%d", uid))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Tạo tên file dựa trên timestamp
	filename := fmt.Sprintf("%d.png", time.Now().UnixNano())
	filePath := filepath.Join(dir, filename)

	// Ghi dữ liệu vào file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
