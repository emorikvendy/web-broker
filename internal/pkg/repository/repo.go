package repository

import "github.com/emorikvendy/web-broker/internal/pkg/model"

type repo interface {
	Get(getReq string) (string, error)
	Put(putReq *model.PutValue) error
}
