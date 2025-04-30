package service

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func TrainCaptcha(uid int32, data []byte, captcha string) (string, string, error) {
	dir := filepath.Join("upload/captcha", fmt.Sprintf("%d", uid))
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

func SaveOtherImage(uid int32, data []byte) (string, error) {
	dir := filepath.Join("upload/other", fmt.Sprintf("%d", uid))
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

// SaveImageWithDedup lưu dữ liệu hình ảnh vào thư mục "upload/images" theo định dạng PNG.
// - Nếu file imgId.png đã tồn tại và nội dung giống nhau: không ghi đè, trả về đường dẫn cũ.
// - Nếu nội dung khác: tạo file mới dạng imgId_2.png, imgId_3.png,... để tránh mất dữ liệu cũ.
// - Trả về đường dẫn tuyệt đối của file vừa được lưu.
//
// Tham số:
//   - uid: hiện chưa dùng, để mở rộng phân vùng ảnh theo người dùng sau này.
//   - imgId: định danh ảnh (dùng làm tên file cơ bản).
//   - data: dữ liệu hình ảnh dạng byte.
//
// Trả về:
//   - Đường dẫn tuyệt đối của file ảnh đã lưu.
//   - isNew: true nếu ảnh được tạo mới, false nếu ảnh đã tồn tại và giống nhau.
//   - error: lỗi nếu có khi tạo thư mục hoặc ghi file.
func SaveImageWithDedup(uid int32, imgId int16, data []byte) (string, bool, error) {
	dir := filepath.Join("upload", "images")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", false, fmt.Errorf("failed to create dir %s: %w", dir, err)
	}
	baseName := fmt.Sprintf("%d", imgId)
	filename := baseName + ".png"
	filePath := filepath.Join(dir, filename)

	if _, err := os.Stat(filePath); err == nil {
		existingData, err := os.ReadFile(filePath)
		if err == nil && bytes.Equal(existingData, data) {
			absPath, _ := filepath.Abs(filePath)
			return absPath, false, nil
		}

		index := 2
		for {
			candidateFile := fmt.Sprintf("%s_%d.png", baseName, index)
			candidatePath := filepath.Join(dir, candidateFile)
			if _, err := os.Stat(candidatePath); os.IsNotExist(err) {
				filePath = candidatePath
				break
			}
			index++
		}
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", false, fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return filePath, true, nil
	}

	return absPath, true, nil
}

// SaveImageIfNotExists lưu ảnh vào thư mục "upload/images" theo tên imgId.png.
// Nếu file đã tồn tại, không ghi lại.
// Trả về đường dẫn tuyệt đối của file và cờ isNew (true nếu file được tạo mới).
func SaveImageIfNotExists(uid int32, imgId int16, data []byte) (string, bool, error) {
	dir := filepath.Join("upload", "images")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", false, fmt.Errorf("failed to create dir %s: %w", dir, err)
	}

	filename := fmt.Sprintf("%d.png", imgId)
	filePath := filepath.Join(dir, filename)

	if _, err := os.Stat(filePath); err == nil {
		absPath, _ := filepath.Abs(filePath)
		return absPath, false, nil
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", false, fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return filePath, true, nil
	}
	return absPath, true, nil
}

func SaveIcon(uid int32, iconId int16, data []byte) (string, bool, error) {
	dir := filepath.Join("upload", "icons")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", false, fmt.Errorf("failed to create dir %s: %w", dir, err)
	}
	baseName := fmt.Sprintf("%d", iconId)
	filename := baseName + ".png"
	filePath := filepath.Join(dir, filename)

	if _, err := os.Stat(filePath); err == nil {
		existingData, err := os.ReadFile(filePath)
		if err == nil && bytes.Equal(existingData, data) {
			absPath, _ := filepath.Abs(filePath)
			return absPath, false, nil
		}

		index := 2
		for {
			candidateFile := fmt.Sprintf("%s_%d.png", baseName, index)
			candidatePath := filepath.Join(dir, candidateFile)
			if _, err := os.Stat(candidatePath); os.IsNotExist(err) {
				filePath = candidatePath
				break
			}
			index++
		}
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", false, fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return filePath, true, nil
	}

	return absPath, true, nil
}
