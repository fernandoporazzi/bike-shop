package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type UploadService interface {
	Upload(formdata *multipart.Form) ([]string, error)
}

type uploadService struct{}

func NewUploadService() UploadService {
	return &uploadService{}
}

func (s *uploadService) Upload(formdata *multipart.Form) ([]string, error) {
	var paths []string
	files := formdata.File["files"]

	for i := range files {
		file, err := files[i].Open()

		defer file.Close()

		if err != nil {
			return paths, err
		}

		// rename file as uuid to avoid naming conflicts
		newName := fmt.Sprintf("%s%s", uuid.NewString(), filepath.Ext(files[i].Filename))

		destination, err := os.Create("./static/" + newName)

		defer destination.Close()

		if err != nil {
			return paths, err
		}

		if _, err := io.Copy(destination, file); err != nil {
			return paths, err
		}

		paths = append(paths, newName)
	}

	return paths, nil
}
