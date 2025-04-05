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

	filenameBase := captcha
	filenameImg := fmt.Sprintf("%s.png", filenameBase)
	// filenameCaptcha := fmt.Sprintf("%s.txt", filenameBase)
	filePathImg := filepath.Join(dir, filenameImg)
	// filePathCaptcha := filepath.Join(dir, filenameCaptcha)

	// Check if files exist and modify the name if necessary
	for i := 1; ; i++ {
		if _, err := os.Stat(filePathImg); os.IsNotExist(err) {
			break
		}
		filenameImg = fmt.Sprintf("%s_%d.png", filenameBase, i)
		// filenameCaptcha = fmt.Sprintf("%s_%d.txt", filenameBase, i)
		filePathImg = filepath.Join(dir, filenameImg)
		// filePathCaptcha = filepath.Join(dir, filenameCaptcha)
	}

	err = os.WriteFile(filePathImg, data, 0644)
	if err != nil {
		return "", "", err
	}

	// err = os.WriteFile(filePathCaptcha, []byte(captcha), 0644)
	// if err != nil {
	// 	return "", "", err
	// }

	// return filePathImg, filePathCaptcha, nil
	return filePathImg, "", nil
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
