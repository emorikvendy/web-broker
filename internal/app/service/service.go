package service

import (
	"github.com/emorikvendy/web-broker/internal/pkg/model"
)

type (
	repo interface {
		Get(key string) (string, error)
		Put(putReq *model.PutValue) error
	}

	Service struct {
		repo repo
	}
)

func New(repo repo) *Service {
	return &Service{repo: repo}
}
