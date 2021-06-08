package controller

import (
	"github.com/yudiwiradinata/go-rest-api-mux/entity"
	"github.com/yudiwiradinata/go-rest-api-mux/service"
)

var (
	postService service.PostService = service.NewPostService()
)

type postController struct{}

type PostController interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

func NewPostController() PostController {
	return &postController{}
}

func (*postController) Validate(post *entity.Post) error {
	return postService.Validate(post)
}

func (*postController) Create(post *entity.Post) (*entity.Post, error) {
	return postService.Create(post)
}

func (*postController) FindAll() ([]entity.Post, error) {
	return postService.FindAll()
}
