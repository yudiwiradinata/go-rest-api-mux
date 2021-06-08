package service

import (
	"errors"

	"github.com/yudiwiradinata/go-rest-api-mux/entity"
	"github.com/yudiwiradinata/go-rest-api-mux/repository"
)

type service struct{}

var (
	postRepository repository.PostRepository
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &service{}
}

func (s *service) FindAll() ([]entity.Post, error) {
	return postRepository.FindAll()
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post, err := postRepository.Save(post)
	if err != nil {
		return nil, errors.New("Error Save Data")
	}
	return post, nil
}

func (s *service) Validate(post *entity.Post) error {
	if post != nil {
		return errors.New("Data post is empty")
	}

	if post.Title == "" {
		return errors.New("Data post title is empty")
	}

	return nil
}
