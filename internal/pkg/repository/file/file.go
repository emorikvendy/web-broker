package file

import (
	"github.com/emorikvendy/web-broker/internal/pkg/model"
)

type FileRepo struct {
	fileName string
}

func New(fileName string) *FileRepo {
	return &FileRepo{fileName: fileName}
}

func (fr *FileRepo) Get(getReq string) (string, error) {
	// TODO: impl
	return "", nil
}

func (fr *FileRepo) Put(putReq *model.PutValue) error {
	// TODO: impl
	return nil
}
