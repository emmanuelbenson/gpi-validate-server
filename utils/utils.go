package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/emmanuelbenson/gpi-validate-v2/random"
)

// Error ...
type Error map[string]interface{}

// Errors ...
var Errors []Error

// Success for registring success message
type Success map[string]interface{}

// WriteFile writes documents or images to Uploads directory
func WriteFile(handle *multipart.FileHeader, uploadPath string) (string, error) {

	genFileName := strings.ToLower(random.String(16))

	// Get file extension
	splitNewFileName := strings.Split(handle.Filename, ".")
	fileExt := splitNewFileName[len(splitNewFileName)-1]

	fileName := genFileName + "." + fileExt

	newPath := filepath.Join(uploadPath, fileName)

	// err := imagewriter.WriteFile(file, newPath)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 	return
	// }

	f, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}
	fz, _ := handle.Open()
	defer f.Close()
	io.Copy(f, fz)
	return fileName, nil
}

// ValidateFileExtension validates both file type
func ValidateFileExtension(handle *multipart.FileHeader) Error {

	splitFilename := strings.Split(handle.Filename, ".")
	fileExt := splitFilename[len(splitFilename)-1]

	if fileExt == "pdf" {
		return nil
	} else if fileExt == "png" {
		return nil
	} else if fileExt == "jpg" {
		return nil
	}
	field := "Document"
	errorMsg := "file type not support"
	err := Error{}
	err[field] = errorMsg

	return err

}

// ValidateFileSize validates file size
func ValidateFileSize(handle *multipart.FileHeader) Error {
	size := handle.Size

	if size > 600000 {
		field := "Document"
		errorMsg := "file size larger than 400KB"
		err := Error{}
		err[field] = errorMsg

		return err
	}
	return nil
}
