package repository

import "github.com/yudiwiradinata/go-rest-api-mux/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
