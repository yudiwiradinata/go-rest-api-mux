package controller

import (
	"errors"

	"github.com/yudiwiradinata/go-rest-api-mux/entity"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type postService struct{}

func NewPostService() PostService {
	return &postService{}
}

func (*postService) Validate(post *entity.Post) error {
	if post != nil {
		return errors.New("Data post is empty")
	}

	if post.Title == "" {
		return errors.New("Data post title is empty")
	}

	return nil
}

func (*postService) Create(post *entity.Post) (*entity.Post, error) {

}

func (*postService) FindAll() ([]entity.Post, error) {

}
